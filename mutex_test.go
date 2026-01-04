package guardedmutex

import (
	"errors"
	"sync"
	"testing"
)

func TestAcquire(t *testing.T) {
	m := Mutex[int]{}

	{
		var captured int
		m.Acquire(func(v int) { captured = v })
		if captured != 0 {
			t.Errorf("expected 0, got %d", captured)
		}

		// Test with a reference type (*int) - modification possible
		// This effectively invalidates the convenience offered by TestAcquireSet, but I'd argue that TestAcquireSet is more readable than using pointers
		val := 10
		mPtr := Mutex[*int]{v: &val}
		mPtr.Acquire(func(v *int) { *v = 20 })
		if val != 20 {
			t.Errorf("expected val to be updated to 20, got %d", val)
		}
	}
}

func TestAcquireErr(t *testing.T) {
	m := Mutex[int]{v: 5}

	{
		// Case 1: No error
		if err := m.AcquireErr(func(v int) error {
			if v != 5 {
				return errors.New("unexpected value")
			}
			return nil
		}); err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Case 2: Error returned
		expectedErr := errors.New("something went wrong")
		if err := m.AcquireErr(func(v int) error { return expectedErr }); err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
	}
}

func TestAcquireSet(t *testing.T) {
	m := Mutex[int]{}

	{
		m.AcquireSet(func(v int) int { return v + 5 })

		m.Acquire(func(v int) {
			if v != 5 {
				t.Errorf("expected 5, got %d", v)
			}
		})
	}
}

func TestConcurrency(t *testing.T) {
	m := Mutex[int]{}

	{
		iterations := 1000
		var wg sync.WaitGroup

		for range iterations {
			wg.Go(func() { m.AcquireSet(func(v int) int { return v + 1 }) })
		}

		wg.Wait()

		m.Acquire(func(v int) {
			if v != iterations {
				t.Errorf("expected %d, got %d", iterations, v)
			}
		})
	}
}
