package internal

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func (session *Session) Start() error {
	var args []string

	if session.Duration > 0 {
		args = append(args, "-t")
		seconds := fmt.Sprintf("%d", int(session.Duration.Seconds()))
		args = append(args, seconds)
	}

	if session.PID != 0 && session.PID != os.Getpid() {
		args = append(args, "-w")
		pid := fmt.Sprintf("%d", session.PID)
		args = append(args, pid)
	}

	caffeinate := exec.Command("caffeinate", args...)
	err := caffeinate.Start()
	if err != nil {
		return err
	}

	exits := make(chan error, 1)
	if session.Duration > 0 {
		go func() {
			err := caffeinate.Wait()
			exits <- err
		}()
	}

	signals := make(chan os.Signal, 1)
	go func() {
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-signals
		err := caffeinate.Process.Kill()
		exits <- err
	}()
	err = <-exits
	return err
}
