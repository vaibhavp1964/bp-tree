package atomiccounter

// AtomicCounter ...
type AtomicCounter struct {
	counter int
}

// NewCounter ...
func NewCounter() *AtomicCounter {
	return &AtomicCounter{
		counter: -1,
	}
}

// NewCounterFromIndex ...
func NewCounterFromIndex(index int) *AtomicCounter {
	return &AtomicCounter{
		counter: index,
	}
}

// Get ...
func (ac *AtomicCounter) Get() int {
	ac.counter++
	return ac.counter
}
