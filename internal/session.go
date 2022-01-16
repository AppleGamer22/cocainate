package internal

import (
	"errors"
	"os"
	"time"
)

type Session struct {
	PID      int
	Duration time.Duration
	active   bool
}

var signals chan os.Signal

// Stop terminates the current session. Can be called only when `session.Wait` is running in the background.
func (session *Session) Stop() error {
	if signals != nil {
		signals <- os.Interrupt
		return nil
	}
	return errors.New("the signal channel has not be initialized, probably because session.Wait is not running in the background")
}
