package assert_test

import (
	"errors"
	"sync"
	"testing"

	"go.flow.arcalot.io/assert/assert"
)

func TestNoError(externalT *testing.T) {
	var internalT1 testing.T // Should fail
	var internalT2 testing.T // Should not fail
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Should fail
		assert.NoError(&internalT1, errors.New("This is an error"))
	}()

	go func() {
		defer wg.Done()
		// Should not fail
		assert.NoError(&internalT2, nil)

	}()

	wg.Wait()

	if !internalT1.Failed() {
		externalT.Fatalf("NoError() should have failed, but didn't.")
	}

	if internalT2.Failed() {
		externalT.Fatalf("NoError() should not have failed, but did.")
	}
}

func TestNoErrorR(externalT *testing.T) {
	var internalT1 testing.T // Should fail
	var internalT2 testing.T // Should not fail
	var wg sync.WaitGroup
	wg.Add(2)
	returned1 := false
	returned2 := false

	go func() {
		defer wg.Done()
		// Should fail
		returned1 = assert.NoErrorR[bool](&internalT1)(true, errors.New("This is an error"))
	}()

	go func() {
		defer wg.Done()
		// Should not fail
		returned2 = assert.NoErrorR[bool](&internalT1)(true, nil)

	}()

	wg.Wait()

	if !internalT1.Failed() {
		externalT.Fatalf("NoErrorR() should have failed, but didn't.")
	}

	if returned1 != false {
		externalT.Fatalf("NoErrorR() didn't properly fail, resulting in a returned value.")
	}

	if internalT2.Failed() {
		externalT.Fatalf("NoErrorR() should not have failed, but did.")
	}

	if returned2 == false {
		externalT.Fatalf("NoErrorR() should have returned a value, but it didn't.")
	}
}

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
