package assert

import (
	"testing"
)

// MapContainsKeyAny asserts that the key is in the map, and returns the value if it passes.
// For use with maps with any as the key type.
func MapContainsKeyAny[V any](t *testing.T, key any, testedMap map[any]V) V {
	t.Helper()
	value, ok := testedMap[key]
	if !ok {
		t.Fatalf("Map does not contain key '%v'", key)
	}
	return value
}

// MapNotContainsKeyAny asserts that the key is not in the map.
// For use with maps with any as the key type.
func MapNotContainsKeyAny[V any](t *testing.T, key any, testedMap map[any]V) {
	t.Helper()
	_, ok := testedMap[key]
	if ok {
		t.Fatalf("Map contains key '%v'", key)
	}
}

// MapContainsKey asserts that the key is in the map, and returns the value if it passes.
// For maps that do not have any as their key type, and instead have a comparable type.
func MapContainsKey[K comparable, V any](t *testing.T, key K, testedMap map[K]V) V {
	t.Helper()
	value, ok := testedMap[key]
	if !ok {
		t.Fatalf("Map does not contain key '%v'", key)
	}
	return value
}

// MapNotContainsKey asserts that the key is not in the map.
// For maps that do not have any as their key type, and instead have a comparable type.
func MapNotContainsKey[K comparable, V any](t *testing.T, key K, testedMap map[K]V) {
	t.Helper()
	_, ok := testedMap[key]
	if ok {
		t.Fatalf("Map contains key '%v'", key)
	}
}
