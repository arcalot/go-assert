package assert_test

import (
    "testing"
)

func testFailure(t *testing.T, f func(t *testing.T)) {
    t.Helper()
    done := make(chan struct{})
    var internalT testing.T
    go func() {
        defer close(done)
        f(&internalT)
    }()
    <-done
    if !internalT.Failed() {
        t.Fatalf("Assertion did not fail.")
    }
}
