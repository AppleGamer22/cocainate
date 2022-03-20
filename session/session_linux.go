package session

import (
	"errors"
	"os"
	"sync"
	"time"

	dbus "github.com/godbus/dbus/v5"
)

const (
	path        = "/org/freedesktop/ScreenSaver"
	screensaver = "org.freedesktop.ScreenSaver"
	inhibit     = "org.freedesktop.ScreenSaver.Inhibit"
	uninhibit   = "org.freedesktop.ScreenSaver.UnInhibit"
)

type Session struct {
	sync.Mutex
	PID      int
	Duration time.Duration
	signals  chan os.Signal
	cookie   uint32
}

/*
Starts the session (according to https://people.freedesktop.org/~hadess/idle-inhibition-spec/re01.html) with a call to the D-BUS screensaver inhibitor.

A non-nil error is returned if the D-BUS session connection fails, if the inhabitation call fails or if the cookie recovery fails.
*/
func (s *Session) Start() error {
	connection, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	defer connection.Close()

	object := connection.Object(screensaver, path)
	call := object.Call(inhibit, 0, "cocainate", "cocainate is running")

	if call.Err != nil {
		return call.Err
	}
	s.Lock()
	defer s.Unlock()
	if err := call.Store(&s.cookie); err != nil {
		return err
	}

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

	connection, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	defer connection.Close()

	s.Lock()
	defer s.Unlock()
	object := connection.Object(screensaver, path)
	err = object.Call(uninhibit, 0, s.cookie).Err
	if err != nil {
		return err
	}

	s.cookie = 0
	return nil
}

// A Boolean for session status
func (s *Session) Active() bool {
	return s.cookie != 0
}
