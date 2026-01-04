package guardedmutex

import (
	"sync"
)

type Mutex[T any] struct {
	mu sync.Mutex
	v  T
}

// Acquire ensures that the mutex lock is guaranteed before calling the provided function with it ensuring safer operations without compromising concurrency safety
func (m *Mutex[T]) Acquire(fn func(v T)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fn(m.v)
}

// AcquireErr is a convenience method that propagates the provided function's error to the caller for error handling at the call site
func (m *Mutex[T]) AcquireErr(fn func(v T) error) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return fn(m.v)
}

// AcquireSet is a convenience method that transforms the acquired value and stores the transformed value for the next read
func (m *Mutex[T]) AcquireSet(fn func(v T) T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.v = fn(m.v)
}
