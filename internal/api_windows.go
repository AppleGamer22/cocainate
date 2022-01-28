package internal

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	esContinuous     = 0x80000000
	esSystemRequired = 0x00000001
)

/*
Starts a SetThreadExecutionState session (https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setthreadexecutionstate).

A non-nil error is returned if the session failed to start.
*/
func (session *Session) Start() error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")
	r1, _, err := setThreadExecStateProc.Call(uintptr(esContinuous | esSystemRequired))
	if r1 == 0 {
		return err
	}

	session.active = true
	return nil
}

/*
Wait can be called only after Start has been called successfully.

Wait will block further execution until the user send an interrupt signal, or until the session duration has passed.

A non-nil error is returned if the SetThreadExecutionState session failed to stop.
*/
func (session *Session) Wait() error {
	if !session.active {
		return errors.New("Wait can be called only after Start has been called successfully")
	}

	exit := make(chan bool, 1)
	if session.Duration > 0 {
		go func() {
			time.Sleep(session.Duration)
			exit <- true
		}()
	}

	if session.Signals == nil {
		session.Signals = make(chan os.Signal, 1)
	}
	go func() {
		signal.Notify(session.Signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-session.Signals
		fmt.Println()
		exit <- true
	}()

	<-exit
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")
	r1, _, err := setThreadExecStateProc.Call(uintptr(esContinuous))
	if r1 == 0 {
		return err
	}

	session.active = false
	return nil
}
