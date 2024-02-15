package assert

import "testing"

// Panics asserts that the function causes a panic.
// Does not detect panics in separate goroutines.
func Panics(t *testing.T, functionToTest func()) {
	defer func() {
		// This deferred function will capture the panic if it panics
		if r := recover(); r == nil {
			t.Fatalf("Expected function to panic, but it didn't.")
		}
	}()

	// Run function to see if it panics.
	functionToTest()
}

// PanicsContains asserts that the function causes a panic.
// It also asserts that the panic's error's string representation has
// the given substring within it.
// Does not detect panics in separate goroutines.
func PanicsContains(t *testing.T, functionToTest func(), expectedErrMsg string) {
	defer func() {
		// This deferred function will capture the panic if it panics
		r := recover()
		if r == nil {
			t.Fatalf("Expected function to panic, but it didn't.")
		}
		err, isErr := r.(error)
		str, isStr := r.(string)
		if !isErr && !isStr {
			t.Fatalf("panic returned value that is neither a string nor an error; got %T", r)
		}
		if !isStr {
			str = err.Error()
		}
		Contains(t, str, expectedErrMsg)
	}()

	// Run function to see if it panics.
	functionToTest()
}

// PanicsWithValidation asserts that the function causes a panic.
// It also calls the function passed in to allow validation of the panicked value.
// Does not detect panics in separate goroutines.
func PanicsWithValidation(
	t *testing.T,
	functionToTest func(),
	validationFunction func(*testing.T, any),
) {
	defer func() {
		// This deferred function will capture the panic if it panics
		r := recover()
		if r == nil {
			t.Fatalf("Expected function to panic, but it didn't.")
		}
		validationFunction(t, r)
	}()

	// Run function to see if it panics.
	functionToTest()
}

// NoPanic asserts that the function does not cause a panic.
// Does not detect panics in separate goroutines.
func NoPanic(t *testing.T, functionToTest func()) {
	defer func() {
		// This deferred function will capture the panic if it panics
		if r := recover(); r != nil {
			t.Fatalf("Expected function to not panic, but it did. Panic output: %v", r)
		}
	}()

	// Run function to see if it panics.
	functionToTest()
}
