package ps

import (
	"errors"
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

/*
Generate a channel for termination signal from an external process (with PID).

A polling interval is used as delay between process checks.
*/
func Notify(pid int32, pollingDuration time.Duration) chan error {
	errs := make(chan error, 1)

	abort := pid == 0 || pid == int32(os.Getpid()) && pollingDuration <= 0

	if abort {
		errs <- errors.New("invalid PID or process polling interval, both must be non-0")
		return errs
	}

	go func() {
		ticker := time.NewTicker(pollingDuration)
		for range ticker.C {
			p, err := process.NewProcess(pid)
			if err != nil {
				errs <- nil
				break
			}

			if running, err := p.IsRunning(); err != nil || !running {
				// if err != nil, a race condition has occurred, the process ended after checking for its existence but before checking if it's running
				errs <- nil
				break
			}
		}
	}()
	return errs
}
