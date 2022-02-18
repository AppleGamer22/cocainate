package session

import (
	"errors"
	"os/exec"
)

/*
Starts a caffeinate (https://github.com/apple-oss-distributions/PowerManagement/tree/main/caffeinate) session.

A non-nil error is returned if the session failed to start.
*/
func (session *Session) Start() error {
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
	defer session.Unlock()
	session.caffeinate = exec.Command("caffeinate")
	if err := session.caffeinate.Start(); err != nil {
		return err
	}

	session.active = true
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
	defer session.Unlock()
	session.active = false
	session.caffeinate = nil
	return nil
}
