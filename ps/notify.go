package ps

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

func Notify(pid int32, pollingDuration time.Duration) chan error {
	fmt.Println(pid)
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
				// process does not exist
				errs <- nil
				break
			}
			running, err := p.IsRunning()
			if err != nil && err.Error() != "exit status 1" {
				errs <- err
				break
			}

			if !running {
				errs <- nil
				break
			}
		}
	}()
	return errs
}
