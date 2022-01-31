package internal

import (
	"errors"
	"os"
	"os/exec"
	"sync"
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
Stop terminates the current session.

Can be called only when Wait is running in the background.
*/
func (session *Session) Stop() error {
	if session.signals == nil {
		return errors.New("the signal channel has not be initialized, probably because session.Wait is not running in the background")
	}

	session.Lock()
	session.signals <- os.Interrupt
	session.Unlock()
	return nil
}
