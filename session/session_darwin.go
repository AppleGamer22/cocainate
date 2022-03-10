package session

import (
	"errors"
	"os/exec"
)

/*
Starts a caffeinate (https://github.com/apple-oss-distributions/PowerManagement/tree/main/caffeinate) session.

A non-nil error is returned if the session failed to start.
*/
func (s *Session) Start() error {
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

	s.Lock()
	defer s.Unlock()
	s.caffeinate = exec.Command("caffeinate")
	if err := s.caffeinate.Start(); err != nil {
		return err
	}

	s.active = true
	return nil
}

/*
Stop kills an already-started session while Wait is not running in the background.

This method is recommended for uses in which the session is required to terminate only by the calling program, and not by the user.
*/
func (s *Session) Stop() error {
	if !s.active || s.caffeinate == nil {
		return errors.New("Wait can be called only after Start has been called successfully")
	}

	if err := s.caffeinate.Process.Kill(); err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()
	s.active = false
	s.caffeinate = nil
	return nil
}
