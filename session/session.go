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

	"github.com/AppleGamer22/cocainate/ps"
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

/*
Creates a New session instance with duration.

If the session's duration is 0, the session will stop after a termination signal or a call to session.Stop.
*/
func New(duration time.Duration, pid int) Session {
	return Session{
		Duration: duration,
		PID:      pid,
		signals:  make(chan os.Signal, 1),
	}
}

/*
Wait can be called only after Start has been called successfully.

Wait will block further execution until the user send an interrupt signal, or until the session duration has passed.

A non-nil error is returned if the D-BUS session connection fails, or if the un-inhabitation call fails.
*/
func (s *Session) Wait() error {
	linuxError := runtime.GOOS == "linux" && (!s.active || s.cookie == 0)
	macError := runtime.GOOS == "darwin" && (!s.active || s.caffeinate == nil)
	windowsError := runtime.GOOS == "windows" && !s.active
	if linuxError || macError || windowsError {
		return errors.New("Wait can be called only after Start has been called successfully")
	}

	if s.signals == nil {
		s.Lock()
		s.signals = make(chan os.Signal, 1)
		s.Unlock()
	}

	signal.Notify(s.signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	if s.Duration > 0 && s.PID != 0 && s.PID != os.Getpid() {
		select {
		case psError := <-ps.Notify(int32(s.PID), s.Duration):
			if stoppingError := s.Stop(); stoppingError != nil && psError != nil {
				return fmt.Errorf("%v\n%v", psError, stoppingError)
			} else {
				return psError
			}
		case <-s.signals:
		}
	} else if s.Duration > 0 {
		select {
		case <-time.After(s.Duration):
		case <-s.signals:
		}
	} else {
		<-s.signals
	}

	return s.Stop()
}

/*
Kill terminates the current session.

Can be called only when Wait is running in the background.
*/
func (s *Session) Kill() error {
	linuxError := runtime.GOOS == "linux" && (!s.active || s.cookie == 0)
	macError := runtime.GOOS == "darwin" && (!s.active || s.caffeinate == nil)
	windowsError := runtime.GOOS == "windows" && !s.active

	if s.signals == nil || linuxError || macError || windowsError {
		return errors.New("Start has not been called successfully or Wait is not running in the background")
	}

	s.Lock()
	s.signals <- os.Interrupt
	s.Unlock()
	return nil
}
