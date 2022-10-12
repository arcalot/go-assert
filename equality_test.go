package assert_test

import (
	"sync"
	"testing"

	"go.arcalot.io/assert"
)

func TestContains(externalT *testing.T) {
	var internalT1 testing.T // Should fail
	var internalT2 testing.T // Should not fail
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Should fail
		assert.Contains(&internalT1, "abc", "d")
	}()

	go func() {
		defer wg.Done()
		// Should not fail
		assert.Contains(&internalT2, "abc", "a")
	}()

	wg.Wait()

	if !internalT1.Failed() {
		externalT.Fatalf("Contains() should have failed, but didn't.")
	}

	if internalT2.Failed() {
		externalT.Fatalf("Contains() should not have failed, but did.")
	}
}

func TestEqualsStr(externalT *testing.T) {
	var internalT1 testing.T // Should fail
	var internalT2 testing.T // Should not fail
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Should fail
		assert.Equals(&internalT1, "a", "b")
	}()

	go func() {
		defer wg.Done()
		// Should not fail
		assert.Equals(&internalT2, "a", "a")
	}()

	wg.Wait()

	if !internalT1.Failed() {
		externalT.Fatalf("Equals() should have failed with str, but didn't.")
	}

	if internalT2.Failed() {
		externalT.Fatalf("Equals() should not have failed with str, but did.")
	}
}

func TestEqualsInt(externalT *testing.T) {
	var internalT1 testing.T // Should fail
	var internalT2 testing.T // Should not fail
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Should fail
		assert.Equals(&internalT1, 1, 2)
	}()

	go func() {
		defer wg.Done()
		// Should not fail
		assert.Equals(&internalT2, 1, 1)
	}()

	wg.Wait()

	if !internalT1.Failed() {
		externalT.Fatalf("Equals() should have failed with int, but didn't.")
	}

	if internalT2.Failed() {
		externalT.Fatalf("Equals() should not have failed with int, but did.")
	}
}

func TestDeepEquals(t *testing.T) {
	var ab1 = map[string]int{
		"a": 1,
		"b": 2,
	}
	var ab2 = map[string]int{
		"a": 1,
		"b": 2,
	}
	var ac = map[string]int{
		"a": 1,
		"c": 3,
	}

	assert.Equals(t, ab1, ab2)
	testFailure(t, func(t *testing.T) {
		assert.Equals(t, ab1, ac)
	})
}
