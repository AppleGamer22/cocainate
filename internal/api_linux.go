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

var cookie uint

// https://people.freedesktop.org/~hadess/idle-inhibition-spec/re01.html
//
// https://youtu.be/-bEzHG2u8XA?t=721
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

	// var cookie uint32
	if err := call.Store(&cookie); err != nil {
		return err
	}

	return nil
}

// Wait can be called only after Start has been called successfully
func (session *Session) Wait() error {
	if cookie == 0 {
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
	return object.Call(uninhibit, 0, cookie).Err
}
