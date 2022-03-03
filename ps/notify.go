package ps

import (
	"errors"
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

func Notify(pid int32, pollingDuration time.Duration) chan error {
	errs := make(chan error, 1)

	abort := pid == 0 || pid == int32(os.Getpid()) && pollingDuration <= 0

	if abort {
		errs <- errors.New("invalid PID or process polling interval, both must be non-0")
		return errs
	}

	go func() {
		p, err := process.NewProcess(pid)
		if err != nil {
			errs <- err
			return
		}

		ticker := time.NewTicker(pollingDuration)
		for range ticker.C {
			running, _ := p.IsRunning()
			if !running {
				errs <- nil
				break
			}
		}
	}()
	return errs
}
