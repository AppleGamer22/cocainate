package session

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

type Session struct {
	sync.Mutex
	PID        int
	Duration   time.Duration
	signals    chan os.Signal
	cookie     uint32
	caffeinate *exec.Cmd
	active     bool
}

func NewSession(pid int, duration time.Duration) Session {
	return Session{
		PID:      pid,
		Duration: duration,
		signals:  make(chan os.Signal, 1),
	}
}

/*
Wait can be called only after Start has been called successfully.

Wait will block further execution until the user send an interrupt signal, or until the session duration has passed.

A non-nil error is returned if the D-BUS session connection fails, or if the un-inhabitation call fails.
*/
func (session *Session) Wait() error {
	linuxError := runtime.GOOS == "linux" && (!session.active || session.cookie == 0)
	macError := runtime.GOOS == "darwin" && (!session.active || session.caffeinate == nil)
	windowsError := runtime.GOOS == "windows" && !session.active
	if linuxError || macError || windowsError {
		return errors.New("Wait can be called only after Start has been called successfully")
	}

	if session.signals == nil {
		session.Lock()
		session.signals = make(chan os.Signal, 1)
		session.Unlock()
	}
	signal.Notify(session.signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	if session.Duration > 0 {
		select {
		case <-time.After(session.Duration):
		case <-session.signals:
			fmt.Print("\b\b")
		}
	} else {
		<-session.signals
		fmt.Print("\b\b")
	}

	return session.Stop()
}

/*
Kill terminates the current session.

Can be called only when Wait is running in the background.
*/
func (session *Session) Kill() error {
	linuxError := runtime.GOOS == "linux" && (!session.active || session.cookie == 0)
	macError := runtime.GOOS == "darwin" && (!session.active || session.caffeinate == nil)
	windowsError := runtime.GOOS == "windows" && !session.active

	if session.signals == nil || linuxError || macError || windowsError {
		return errors.New("Start has not been called successfully or Wait is not running in the background")
	}

	session.Lock()
	session.signals <- os.Interrupt
	session.Unlock()
	return nil
}
