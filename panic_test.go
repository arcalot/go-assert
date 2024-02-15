package assert_test

import (
	"fmt"
	"testing"

	"go.arcalot.io/assert"
)

func panicsTest() {
	panic("This is for testing purposes")
}

func panicWithArgs(panicValue any, _ int) {
	panic(panicValue)
}

func TestCatchesPanic(t *testing.T) {
	// Does panic
	assert.Panics(t, panicsTest)
	assert.Panics(t, func() {
		panicWithArgs("This is for testing purposes", 0)
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

func TestCatchesPanicString(t *testing.T) {
	// Does panic
	// Test with string panic
	assert.PanicsContains(
		t,
		func() {
			panicWithArgs("This is for testing purposes", 0)
		},
		"testing purposes",
	)
	// Test with error panic
	assert.PanicsContains(
		t,
		func() {
			panicWithArgs(fmt.Errorf("this is for testing purposes"), 0)
		},
		"testing purposes",
	)
	// Panic value fails
	// Test with string
	testFailure(t, func(t *testing.T) {
		assert.PanicsContains(
			t,
			func() {
				panicWithArgs("This is for testing purposes", 0)
			},
			"wrong substr",
		)
	})
	// Test with error
	testFailure(t, func(t *testing.T) {
		assert.PanicsContains(
			t,
			func() {
				panicWithArgs(fmt.Errorf("this is for testing purposes"), 0)
			},
			"wrong substr",
		)
	})
	// Test incompatible panic. PanicsContains expects a string or error type.
	testFailure(t, func(t *testing.T) {
		assert.PanicsContains(
			t,
			func() {
				panicWithArgs(0, 0)
			},
			"abc",
		)
	})
	// Does not panic
	testFailure(t, func(t *testing.T) {
		assert.PanicsContains(
			t,
			func() {},
			"",
		)
	})
}

func TestCatchesPanicWithValidation(t *testing.T) {
	// Does panic
	// Test with int panic
	assert.PanicsWithValidation(
		t,
		func() {
			panicWithArgs(1, 0)
		},
		func(t *testing.T, panicValue any) {
			assert.Equals(t, panicValue, 1)
		},
	)
	// Test with string panic
	assert.PanicsWithValidation(
		t,
		func() {
			panicWithArgs("abc", 0)
		},
		func(t *testing.T, panicValue any) {
			assert.Equals(t, panicValue, "abc")
		},
	)
	// Panic's value validation fails
	// Test with string
	testFailure(t, func(t *testing.T) {
		assert.PanicsWithValidation(
			t,
			func() {
				panicWithArgs("This is for testing purposes", 0)
			},
			func(t *testing.T, panicValue any) {
				assert.Equals(t, panicValue, "abc")
			},
		)
	})
	// Does not panic
	testFailure(t, func(t *testing.T) {
		assert.PanicsWithValidation(
			t,
			func() {},
			func(t *testing.T, a any) {},
		)
	})
}
