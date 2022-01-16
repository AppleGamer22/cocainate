package internal

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	esContinuous     = 0x80000000
	esSystemRequired = 0x00000001
)

// Start
//
// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-powersetrequest
//
// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setthreadexecutionstate
//
// https://stackoverflow.com/questions/45436158/how-to-to-stop-a-machine-from-sleeping-hibernating-for-execution-period
//
// https://github.com/iamacarpet/go-win64api/blob/master/process.go#L143-L150
func (session *Session) Start() error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")
	r1, _, err := setThreadExecStateProc.Call(uintptr(esContinuous | esSystemRequired))
	if r1 == 0 {
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
		fmt.Println()
		exit <- true
	}()

	<-exit
	r1, _, err = setThreadExecStateProc.Call(uintptr(esContinuous))
	if r1 == 0 {
		return err
	}
	return nil
}
