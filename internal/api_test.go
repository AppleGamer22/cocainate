package internal_test

import (
	"sync"
	"testing"
	"time"

	"github.com/AppleGamer22/cocainate/internal"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test for session duration
func TestDuration(t *testing.T) {
	session := internal.NewSession(0, time.Nanosecond)
	err := session.Start()
	require.NoError(t, err)

	err = session.Wait()
	require.NoError(t, err)
}

// Test for session interrupt signal
func TestInterrupt(t *testing.T) {
	session := internal.NewSession(0, 0)
	err := session.Start()
	require.NoError(t, err)

	err = session.Kill()
	require.NoError(t, err)

	err = session.Wait()
	require.NoError(t, err)
}

// Test for session programtic stop while Wait is running
func TestKill(t *testing.T) {
	session := internal.NewSession(0, 0)
	err := session.Start()
	require.NoError(t, err)

	errs := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		errs <- session.Wait()
		wg.Done()
	}()

	go func() {
		errs <- session.Kill()
		wg.Done()
	}()

	wg.Wait()
	for i := 0; i < cap(errs); i++ {
		err := <-errs
		require.NoError(t, err)
	}
}

// Test for session programtic stop while Wait is not running
func TestStop(t *testing.T) {
	session := internal.NewSession(0, 0)
	err := session.Start()
	require.NoError(t, err)

	err = session.Stop()
	require.NoError(t, err)
}

// Test for when Wait is called before Start
func TestWaitBeforeStart(t *testing.T) {
	session := internal.NewSession(0, 0)
	err := session.Wait()
	require.Error(t, err)
}
