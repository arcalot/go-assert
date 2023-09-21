package assert

import (
	"reflect"
	"testing"
)

// SliceContains asserts that the value is in the list. Returns the index. O log N runtime.
func SliceContains[V any](t *testing.T, expected any, testedSlice []V) int {
	t.Helper()
	for i, value := range testedSlice {
		if reflect.DeepEqual(value, expected) {
			return i
		}
	}
	t.Fatalf("Slice does not contain value '%v'", expected)
	return -1
}

// SliceNotContains asserts that the value is not in the list. O log N runtime.
func SliceNotContains[V any](t *testing.T, unexpectedValue any, testedSlice []V) {
	t.Helper()
	for _, value := range testedSlice {
		if reflect.DeepEqual(value, unexpectedValue) {
			t.Fatalf("Slice contains value '%v'", unexpectedValue)
		}
	}
}

// SliceContainsMatch checks each value with your test function.
// The test function should take the type of one item in the list, and return true if the expected value is found.
func SliceContainsMatch[V any](t *testing.T, testFunction func(V) bool, testedSlice []V) int {
	t.Helper()
	for i, value := range testedSlice {
		if testFunction(value) {
			return i
		}
	}
	t.Fatalf("Slice does not contain a value that satisfies the given testFunction'")
	return -1
}

// SliceContainsExtractor allows you to pass in a function to extract a value from the given type to compare.
// Returns the index of the value.
func SliceContainsExtractor[T any, V any](t *testing.T, valueExtractor func(T) V, expected V, testedSlice []T) int {
	t.Helper()
	for i, value := range testedSlice {
		if reflect.DeepEqual(expected, valueExtractor(value)) {
			return i
		}
	}
	t.Fatalf("Slice does not contain an extracted value that matches the expected value '%v'.'", expected)
	return -1
}
