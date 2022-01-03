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
func Start(duration time.Duration, pid int) error {
	connection, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	defer connection.Close()
	object := connection.Object(screensaver, path)
	call := object.Call(inhibit, 0, "cocainate", "cocainate is running")
	var cookie uint32
	if err := call.Store(&cookie); err != nil {
		return err
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	call = object.Call(uninhibit, 0, cookie)
	return call.Err
}
