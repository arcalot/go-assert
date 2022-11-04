package assert_test

import (
	"testing"

	"go.arcalot.io/assert"
)

func panicsTest() {
	panic("This is for testing purposes")
}

func panicWithArgs(_ int, _ int) {
	panic("This is for testing purposes")
}

func TestCatchesPanic(t *testing.T) {
	// Does panic
	assert.Panics(t, panicsTest)
	assert.Panics(t, func() {
		panicWithArgs(0, 0)
	})
	// Does not panic
	testFailure(t, func(t *testing.T) {
		assert.Panics(t, func() {})
	})
}

func TestNoPanic(t *testing.T) {
	// Does not panic
	assert.NoPanic(t, func() {})

	// Does panic
	testFailure(t, func(t *testing.T) {
		assert.NoPanic(t, panicsTest)
	})
	// Does panic
	testFailure(t, func(t *testing.T) {
		assert.NoPanic(t, func() {
			panicWithArgs(0, 0)
		})
	})
}
