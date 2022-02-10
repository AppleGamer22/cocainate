package internal

import (
	"errors"
	"os/exec"
)

/*
Starts a caffeinate (https://github.com/apple-oss-distributions/PowerManagement/tree/main/caffeinate) session.

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
Stop kills an already-started session while Wait is not running in the background.

This method is recommended for uses in which the session is required to terminate only by the calling program, and not by the user.
*/
func (session *Session) Stop() error {
	if !session.active || session.caffeinate == nil {
		return errors.New("Wait can be called only after Start has been called successfully")
	}

	if err := session.caffeinate.Process.Kill(); err != nil {
		return err
	}

	session.Lock()
	session.active = false
	session.caffeinate = nil
	session.Unlock()
	return nil
}
