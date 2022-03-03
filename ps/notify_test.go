package ps_test

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/AppleGamer22/cocainate/ps"
	"github.com/stretchr/testify/require"
)

func TestNotify(t *testing.T) {
	process := exec.Command("man", "ls")
	err := process.Start()
	require.NoError(t, err)
	pid := process.Process.Pid
	fmt.Println(pid)
	go process.Process.Kill()
	err = <-ps.Notify(int32(pid), time.Second)
	require.NoError(t, err)
}
