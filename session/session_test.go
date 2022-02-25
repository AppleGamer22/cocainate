package session_test

import (
	"sync"
	"testing"
	"time"

	"github.com/AppleGamer22/cocainate/session"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test for session duration
func TestDuration(t *testing.T) {
	s := session.New(time.Nanosecond)
	err := s.Start()
	require.NoError(t, err)

	err = s.Wait()
	require.NoError(t, err)
}

// Test for session interrupt signal
func TestInterrupt(t *testing.T) {
	s := session.New(0)
	err := s.Start()
	require.NoError(t, err)

	err = s.Kill()
	require.NoError(t, err)

	err = s.Wait()
	require.NoError(t, err)
}

// Test for session programmatic stop while Wait is running
func TestKill(t *testing.T) {
	s := session.New(0)
	err := s.Start()
	require.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err := s.Wait()
		require.NoError(t, err)
		wg.Done()
	}()

	go func() {
		err := s.Kill()
		require.NoError(t, err)
		wg.Done()
	}()

	wg.Wait()
}

// Test for session programmatic stop while Wait is not running
func TestStop(t *testing.T) {
	s := session.New(0)
	err := s.Start()
	require.NoError(t, err)

	err = s.Stop()
	require.NoError(t, err)
}

// Test for when Wait is called before Start
func TestErrors(t *testing.T) {
	s := session.New(0)
	err := s.Wait()
	require.Error(t, err)

	err = s.Kill()
	require.Error(t, err)
}
