package internal

import (
	// "fmt"
	"context"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func (session *Session) Start() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	defer func() {
		signal.Stop(signals)
		cancel()
	}()
	// t := ""
	// if duration > 0 {
	// 	t = fmt.Sprintf("-t %s", duration)
	// }
	// w := ""
	// if pid > 0 {
	// 	w = fmt.Sprintf("-w %d", pid)
	// }
	cmd := exec.Command("caffeinate")
	err := cmd.Run()
	if err.Error() == "signal: interrupt" {
		return nil
	}
	return err
}
