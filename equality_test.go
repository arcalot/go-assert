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

type TestStructA struct {
	a int
}

func TestInstanceOf(t *testing.T) {
	var isInt any = 1
	var isFloat any = 1.0
	var isTestStructA any = TestStructA{1}
	var isTestStructAPtr any = &TestStructA{1}

	assert.InstanceOf[int](t, isInt)
	assert.InstanceOf[float64](t, isFloat)
	assert.InstanceOf[TestStructA](t, isTestStructA)
	assert.InstanceOf[*TestStructA](t, isTestStructAPtr)
	assert.InstanceOf[any](t, isTestStructAPtr)

	testFailure(t, func(t *testing.T) {
		assert.InstanceOf[float64](t, isInt)
	})
	testFailure(t, func(t *testing.T) {
		assert.InstanceOf[int](t, isFloat)
	})
	testFailure(t, func(t *testing.T) {
		assert.InstanceOf[TestStructA](t, isTestStructAPtr)
	})
	testFailure(t, func(t *testing.T) {
		assert.InstanceOf[*TestStructA](t, isTestStructA)
	})
	testFailure(t, func(t *testing.T) {
		assert.InstanceOf[int](t, isTestStructA)
	})
	testFailure(t, func(t *testing.T) {
		assert.InstanceOf[int](t, isTestStructAPtr)
	})
}

func TestGreaterThan(t *testing.T) {
	assert.GreaterThan(t, 6, 5)
	assert.GreaterThan(t, 60000000, 0)
	assert.GreaterThan(t, 6, -5)
	assert.GreaterThan(t, 6.0, 5.0)
	assert.GreaterThan(t, "b", "a")

	testFailure(t, func(t *testing.T) {
		assert.GreaterThan(t, 0, 0)
	})
	testFailure(t, func(t *testing.T) {
		assert.GreaterThan(t, 5, 6)
	})
	testFailure(t, func(t *testing.T) {
		assert.GreaterThan(t, "a", "b")
	})
}

func TestLessThan(t *testing.T) {
	assert.LessThan(t, 5, 6)
	assert.LessThan(t, 0, 6000000)
	assert.LessThan(t, -5, 6)
	assert.LessThan(t, 5.0, 6.0)
	assert.LessThan(t, "a", "b")

	testFailure(t, func(t *testing.T) {
		assert.LessThan(t, 0, 0)
	})
	testFailure(t, func(t *testing.T) {
		assert.LessThan(t, 6, 5)
	})
	testFailure(t, func(t *testing.T) {
		assert.LessThan(t, "b", "a")
	})
}

func TestGreaterThanEq(t *testing.T) {
	assert.GreaterThanEq(t, 1, 1)
	assert.GreaterThanEq(t, 6, 5)
	assert.GreaterThanEq(t, 60000000, 0)
	assert.GreaterThanEq(t, 6, -5)
	assert.GreaterThanEq(t, 6.0, 5.0)
	assert.GreaterThanEq(t, "b", "a")
	assert.GreaterThanEq(t, "a", "a")

	testFailure(t, func(t *testing.T) {
		assert.GreaterThanEq(t, 5, 6)
	})
	testFailure(t, func(t *testing.T) {
		assert.GreaterThanEq(t, "a", "b")
	})
}

func TestLessThanEq(t *testing.T) {
	assert.LessThanEq(t, 5, 6)
	assert.LessThanEq(t, 0, 6000000)
	assert.LessThanEq(t, -5, 6)
	assert.LessThanEq(t, 5.0, 6.0)
	assert.LessThanEq(t, "a", "b")
	assert.LessThanEq(t, "a", "a")
	assert.LessThanEq(t, 0, 0)

	testFailure(t, func(t *testing.T) {
		assert.LessThanEq(t, 6, 5)
	})
	testFailure(t, func(t *testing.T) {
		assert.LessThanEq(t, "b", "a")
	})
}
