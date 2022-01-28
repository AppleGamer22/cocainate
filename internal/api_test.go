package internal

import (
	"os"
	"testing"
	"time"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test for session duration
func TestDuration(t *testing.T) {
	session := Session{
		Duration: time.Nanosecond,
	}

	err := session.Start()
	require.NoError(t, err)

	err = session.Wait()
	require.NoError(t, err)
}

// Test for session interrupt
func TestInterrupt(t *testing.T) {
	session := Session{}
	err := session.Start()
	require.NoError(t, err)

	session.Signals = make(chan os.Signal, 1)
	session.Signals <- os.Interrupt

	err = session.Wait()
	require.NoError(t, err)
}
