package ps

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

func Notify(pid int32, pollingDuration time.Duration, signals chan os.Signal) {
	go func() {
		ticker := time.NewTicker(pollingDuration)
		for range ticker.C {
			exists, err := process.PidExists(pid)
			if err != nil || !exists {
				signals <- os.Interrupt
				break
			}
		}
	}()
}
