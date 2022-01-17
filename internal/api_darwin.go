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

var caffeinate *exec.Cmd

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

	caffeinate = exec.Command("caffeinate", args...)
	return caffeinate.Start()
}

// Wait can be called only after Start has been called successfully
func (session *Session) Wait() error {
	if caffeinate == nil {
		return errors.New("Wait can be called only after Start has been called successfully")
	}

	exits := make(chan error, 1)
	if session.Duration > 0 {
		go func() {
			time.Sleep(session.Duration)
			err := caffeinate.Process.Kill()
			exits <- err
		}()
	}

	if session.Signals == nil {
		session.Signals = make(chan os.Signal, 1)
	}
	go func() {
		signal.Notify(session.Signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-session.Signals
		fmt.Println()
		err := caffeinate.Process.Kill()
		exits <- err
	}()
	err := <-exits
	return err
}
