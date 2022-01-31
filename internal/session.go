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
	Signals    chan os.Signal
	cookie     uint32
	caffeinate *exec.Cmd
	active     bool
}

/*
Stop terminates the current session.

Can be called only when Wait is running in the background.
*/
func (session *Session) Stop() error {
	if session.Signals != nil {
		session.Signals <- os.Interrupt
		return nil
	}
	return errors.New("the signal channel has not be initialized, probably because session.Wait is not running in the background")
}
