package session

import (
	"errors"
	"os"
	"sync"
	"syscall"
	"time"
)

const (
	esContinuous     = 0x80000000
	esSystemRequired = 0x00000001
)

type Session struct {
	sync.Mutex
	PID      int
	Duration time.Duration
	signals  chan os.Signal
	active   bool
}

/*
Starts a SetThreadExecutionState session (https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setthreadexecutionstate).

A non-nil error is returned if the session failed to start.
*/
func (s *Session) Start() error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")
	r1, _, err := setThreadExecStateProc.Call(uintptr(esContinuous | esSystemRequired))
	if r1 == 0 {
		return err
	}

	s.Lock()
	s.active = true
	s.Unlock()
	return nil
}

/*
Stop kills an already-started session while Wait is not running in the background.

This method is recommended for uses in which the session is required to terminate only by the calling program, and not by the user.
*/
func (s *Session) Stop() error {
	if !s.Active() {
		return errors.New("Stop can be called only after Start has been called successfully")
	}

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")
	r1, _, err := setThreadExecStateProc.Call(uintptr(esContinuous))
	if r1 == 0 {
		return err
	}

	s.Lock()
	s.active = false
	s.Unlock()
	return nil
}

// A Boolean for session status
func (s *Session) Active() bool {
	return s.active
}
