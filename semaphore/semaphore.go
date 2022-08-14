package semaphore

import "sync"

// Semaphore is a semaphore implementation.
type Semaphore struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

// NewSemaphore creates a new semaphore with the given limit.
func NewSemaphore(maxConcurrentOps uint) *Semaphore {
	return &Semaphore{make(chan struct{}, maxConcurrentOps), new(sync.WaitGroup)}
}

// Acquire acquires a semaphore.
func (s *Semaphore) Acquire() {
	s.wg.Add(1)
	s.c <- struct{}{}
}

// Release releases a semaphore.
func (s *Semaphore) Release() {
	<-s.c
	s.wg.Done()
}

// Wait waits for all acquired semaphores to be released.
func (s *Semaphore) Wait() {
	s.wg.Wait()
}

// Limit returns the limit of the semaphore.
func (s *Semaphore) Limit() int {
	return cap(s.c)
}

// Count returns the number of acquired semaphores.
func (s *Semaphore) Count() int {
	return len(s.c)
}
