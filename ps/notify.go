package ps

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

func Notify(pid int32, pollingDuration time.Duration) chan error {
	signals := make(chan error, 1)

	abort := pid == 0 || pid == int32(os.Getpid()) && pollingDuration <= 0

	if abort {
		return signals
	}

	go func() {
		p, err := process.NewProcess(pid)
		if err != nil {
			signals <- err
			return
		}

		ticker := time.NewTicker(pollingDuration)
		for range ticker.C {
			running, _ := p.IsRunning()
			if !running {
				signals <- nil
				break
			}
		}
	}()
	return signals
}
