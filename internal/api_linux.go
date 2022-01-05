package internal

import (
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

// https://people.freedesktop.org/~hadess/idle-inhibition-spec/re01.html
//
// https://youtu.be/-bEzHG2u8XA
func (session *Session) Start() error {
	connection, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	defer connection.Close()

	object := connection.Object(screensaver, path)
	call := object.Call(inhibit, 0, "cocainate", "cocainate is running")
	var cookie uint32
	if err := call.Store(&cookie); err != nil || call.Err != nil {
		return err
	}

	exit := make(chan bool, 1)
	if session.Duration > 0 {
		go func() {
			time.Sleep(session.Duration)
			exit <- true
		}()
	}

	signals := make(chan os.Signal, 1)
	go func() {
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-signals
		exit <- true
	}()

	<-exit
	return object.Call(uninhibit, 0, cookie).Err
}
