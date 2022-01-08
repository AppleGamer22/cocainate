package internal

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	ESContinuous     = 0x80000000
	ESSystemRequired = 0x00000001
)

// Start
//
// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-powersetrequest
//
// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setthreadexecutionstate
//
// https://stackoverflow.com/questions/45436158/how-to-to-stop-a-machine-from-sleeping-hibernating-for-execution-period
//
// https://github.com/iamacarpet/go-win64api/blob/master/process.go
func (session *Session) Start() error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")
	_, _, err := setThreadExecStateProc.Call(uintptr(ESContinuous | ESSystemRequired))
	if err != nil && !strings.Contains(err.Error(), "operation completed successfully") {
		return err
	}

	exit := make(chan bool, 1)
	if session.Duration > 0 {
		go func() {
			time.Sleep(session.Duration)
			exit <- true
		}()
	}

	signals := make(chan os.Signal, 1)
	go func() {
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-signals
		exit <- true
	}()

	<-exit
	_, _, err = setThreadExecStateProc.Call(uintptr(ESContinuous))
	if err != nil && !strings.Contains(err.Error(), "operation completed successfully") {
		return err
	}
	return nil
}
