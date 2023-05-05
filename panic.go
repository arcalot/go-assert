package assert

import "testing"

// Panics asserts that the function causes a panic.
// Does not detect panics in separate goroutines.
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

// NoPanic asserts that the function does not cause a panic.
// Does not detect panics in separate goroutines.
func NoPanic(t *testing.T, functionToTest func()) {
	defer func() {
		// This defered function will recover it if it panics
		if r := recover(); r != nil {
			t.Fatalf("Expected function to not panic, but it did. Panic output: %v", r)
		}
	}()

	// Run function to see if it panics.
	functionToTest()
}
