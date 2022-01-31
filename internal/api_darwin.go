package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

/*
Starts a caffeinate (https://github.com/apple-oss-distributions/PowerManagement/tree/PowerManagement-1132.141.1/caffeinate) session.

A non-nil error is returned if the session failed to start.
*/
func (session *Session) Start() error {
	var args []string

	// if session.Duration > 0 {
	// 	args = append(args, "-t")
	// 	seconds := fmt.Sprintf("%d", int(session.Duration.Round(time.Second)))
	// 	args = append(args, seconds)
	// }

	// if session.PID != 0 && session.PID != os.Getpid() {
	// 	args = append(args, "-w")
	// 	pid := fmt.Sprintf("%d", session.PID)
	// 	args = append(args, pid)
	// }

	session.Lock()
	session.caffeinate = exec.Command("caffeinate", args...)
	err := session.caffeinate.Start()
	if err != nil {
		return nil
	}

	session.active = true
	session.Unlock()
	return nil
}

/*
Wait can be called only after Start has been called successfully.

Wait will block further execution until the user send an interrupt signal, or until the session duration has passed.

A non-nil error is returned if the caffeinate session failed to stop.
*/
func (session *Session) Wait() error {
	if !session.active || session.caffeinate == nil {
		return errors.New("Wait can be called only after Start has been called successfully")
	}

	exits := make(chan error, 1)
	if session.Duration > 0 {
		go func() {
			time.Sleep(session.Duration)
			err := session.caffeinate.Process.Kill()
			exits <- err
		}()
	}

	if session.Signals == nil {
		session.Signals = make(chan os.Signal, 1)
	}
	go func() {
		signal.Notify(session.Signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-session.Signals
		fmt.Print("\b\b")
		err := session.caffeinate.Process.Kill()
		exits <- err
	}()
	err := <-exits
	if err != nil {
		return err
	}

	session.Lock()
	session.active = false
	session.caffeinate = nil
	session.Unlock()
	return nil
}
