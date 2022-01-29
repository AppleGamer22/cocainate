package internal

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	dbus "github.com/godbus/dbus/v5"
)

const (
	path        = "/org/freedesktop/ScreenSaver"
	screensaver = "org.freedesktop.ScreenSaver"
	inhibit     = "org.freedesktop.ScreenSaver.Inhibit"
	uninhibit   = "org.freedesktop.ScreenSaver.UnInhibit"
)

/*
Starts the session (according to https://people.freedesktop.org/~hadess/idle-inhibition-spec/re01.html) with a call to the D-BUS screensaver inhibitor.

A non-nil error is returned if the D-BUS session connection fails, if the inhabitation call fails or if the cookie recovery fails.
*/
func (session *Session) Start() error {
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

	if err := call.Store(&session.cookie); err != nil {
		return err
	}

	session.active = true
	return nil
}

/*
Wait can be called only after Start has been called successfully.

Wait will block further execution until the user send an interrupt signal, or until the session duration has passed.

A non-nil error is returned if the D-BUS session connection fails, or if the uninhabitation call fails.
*/
func (session *Session) Wait() error {
	if !session.active || session.cookie == 0 {
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
	connection, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	defer connection.Close()

	object := connection.Object(screensaver, path)
	err = object.Call(uninhibit, 0, session.cookie).Err
	if err != nil {
		return err
	}

	session.active = false
	session.cookie = 0
	return nil
}
