package ps_test

import (
	"os/exec"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/AppleGamer22/cocainate/ps"
	"github.com/stretchr/testify/assert"
)

func TestNotify(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	cmd := func() *exec.Cmd {
		if runtime.GOOS != "windows" {
			return exec.Command("man", "man")
		}
		return exec.Command("help")
	}()

	err := cmd.Start()
	assert.NoError(t, err)
	assert.NotNil(t, cmd.Process)
	pid := cmd.Process.Pid

	go func() {
		err := cmd.Wait()
		assert.Error(t, err)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Nanosecond)
		err := <-ps.Notify(int32(pid), time.Nanosecond)
		assert.NoError(t, err)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Nanosecond * 2)
		err := cmd.Process.Kill()
		assert.NoError(t, err)
		wg.Done()
	}()

	wg.Wait()
}
