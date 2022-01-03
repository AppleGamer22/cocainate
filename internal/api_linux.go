package internal

import (
	"os"
	"os/signal"
	"time"

	"github.com/godbus/dbus/v5"
)

const (
	screenSaverPath = "org.freedesktop.ScreenSaver"
	inhibit         = "org.freedesktop.ScreenSaver.Inhibit"
	uninhibit       = "org.freedesktop.ScreenSaver.UnInhibit"
)

// https://people.freedesktop.org/~hadess/idle-inhibition-spec/re01.html
func Start(duration time.Duration, pid int) error {
	connection, err := dbus.ConnectSessionBus()
	if err != nil {
		return err
	}
	defer connection.Close()
	// headers := make(map[dbus.HeaderField]dbus.Variant)
	// headers[dbus.FieldDestination] = dbus.MakeVariant(inhibit)
	// // headers[dbus.FieldInterface] = dbus.MakeVariant()
	// message := dbus.Message{}
	// message.Type = dbus.TypeMethodCall
	// message.Headers = headers

	// if err := connection.Emit(screenSaverPath, inhibit, "cocainate", "cocainate is running"); err != nil {
	// 	return err
	// }
	call1 := connection.BusObject().Call(inhibit, dbus.FlagAllowInteractiveAuthorization, "cocainate", "cocainate is running")
	if call1.Err != nil {
		return err
	}
	cookie := call1.Body
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	call2 := connection.BusObject().Call(uninhibit, dbus.FlagAllowInteractiveAuthorization, cookie...)
	if call2.Err != nil {
		return err
	}
	return nil
}
