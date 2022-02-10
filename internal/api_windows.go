package internal

import (
	"errors"
	"syscall"
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

	session.Lock()
	session.active = true
	session.Unlock()
	return nil
}

/*
Stop kills an already-started session while Wait is not running in the background.

This method is recommended for uses in which the session is required to terminate only by the calling program, and not by the user.
*/
func (session *Session) Stop() error {
	if !session.active {
		return errors.New("Stop can be called only after Start has been called successfully")
	}

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")
	r1, _, err := setThreadExecStateProc.Call(uintptr(esContinuous))
	if r1 == 0 {
		return err
	}

	session.Lock()
	session.active = false
	session.Unlock()
	return nil
}
