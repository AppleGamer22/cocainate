package ps_test

import (
	"os/exec"
	"sync"
	"testing"
	"time"

	"github.com/AppleGamer22/cocainate/ps"
	"github.com/shirou/gopsutil/process"
	"github.com/stretchr/testify/require"
)

func TestNotify(t *testing.T) {
	cmd := exec.Command("man", "ls")
	err := cmd.Start()
	require.NoError(t, err)
	pid := cmd.Process.Pid
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err = <-ps.Notify(int32(pid), time.Nanosecond)
		require.NoError(t, err)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Nanosecond * 2)
		p, err := process.NewProcess(int32(pid))
		require.NoError(t, err)
		err = p.Kill()
		require.NoError(t, err)
		wg.Done()
	}()

	wg.Wait()
}
