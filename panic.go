package assert

import "testing"

// Panics asserts that the function causes a panic.
func Panics(t *testing.T, functionToTest func()) {
	defer func() {
		// This defered function will recover it if it panics
		if r := recover(); r == nil {
			t.Fatalf("Expected function to panic, but it didn't.")
		}
	}()

	// Run function to see if it panics.
	functionToTest()
}

// Panics asserts that the function causes a panic.
func NoPanic(t *testing.T, functionToTest func()) {
	defer func() {
		// This defered function will recover it if it panics
		if r := recover(); r != nil {
			t.Fatalf("Expected function to not panic, but it did.")
		}
	}()

	// Run function to see if it panics.
	functionToTest()
}
