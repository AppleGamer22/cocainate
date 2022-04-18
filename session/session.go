package session

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AppleGamer22/cocainate/ps"
)

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

A non-nil error is returned if the un-inhabitation call fails.
*/
func (s *Session) Wait() error {
	if !s.Active() {
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
	if s.signals == nil || !s.Active() {
		return errors.New("Start has not been called successfully or Wait is not running in the background")
	}

	s.Lock()
	s.signals <- os.Interrupt
	s.Unlock()
	return nil
}
