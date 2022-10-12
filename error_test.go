package assert_test

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"go.arcalot.io/assert"
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

func TestError(externalT *testing.T) {
	var internalT1 testing.T // Should fail
	var internalT2 testing.T // Should not fail
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Should fail
		assert.Error(&internalT1, nil)
	}()

	go func() {
		defer wg.Done()
		// Should not fail
		assert.Error(&internalT2, fmt.Errorf("this is an error"))

	}()

	wg.Wait()

	if !internalT1.Failed() {
		externalT.Fatalf("Error() should have failed, but didn't.")
	}

	if internalT2.Failed() {
		externalT.Fatalf("Error() should not have failed, but did.")
	}
}

func TestErrorR(externalT *testing.T) {
	var internalT1 testing.T // Should fail
	var internalT2 testing.T // Should not fail
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Should fail
		assert.ErrorR[bool](&internalT1)(true, nil)
	}()

	go func() {
		defer wg.Done()
		// Should not fail
		assert.ErrorR[bool](&internalT1)(true, errors.New("This is an error"))
	}()

	wg.Wait()

	if !internalT1.Failed() {
		externalT.Fatalf("ErrorR() should have failed, but didn't.")
	}

	if internalT2.Failed() {
		externalT.Fatalf("ErrorR() should not have failed, but did.")
	}
}
