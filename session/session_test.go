package session_test

import (
	"sync"
	"testing"
	"time"

	"github.com/AppleGamer22/cocainate/session"

	"github.com/stretchr/testify/assert"
)

// Test for session duration
func TestDuration(t *testing.T) {
	s := session.New(time.Nanosecond, 0)
	err := s.Start()
	assert.NoError(t, err)

	err = s.Wait()
	assert.NoError(t, err)
}

// Test for session interrupt signal
func TestInterrupt(t *testing.T) {
	s := session.New(0, 0)
	err := s.Start()
	assert.NoError(t, err)

	err = s.Kill()
	assert.NoError(t, err)

	err = s.Wait()
	assert.NoError(t, err)
}

// Test for session programmatic stop while Wait is running
func TestKill(t *testing.T) {
	s := session.New(0, 0)
	err := s.Start()
	assert.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err := s.Wait()
		assert.NoError(t, err)
		wg.Done()
	}()

	go func() {
		err := s.Kill()
		assert.NoError(t, err)
		wg.Done()
	}()

	wg.Wait()
}

// Test for session programmatic stop while Wait is not running
func TestStop(t *testing.T) {
	s := session.New(0, 0)
	err := s.Start()
	assert.NoError(t, err)

	err = s.Stop()
	assert.NoError(t, err)
}

// Test for when Wait is called before Start
func TestErrors(t *testing.T) {
	s := session.New(0, 0)
	err := s.Wait()
	assert.Error(t, err)

	err = s.Kill()
	assert.Error(t, err)
}
