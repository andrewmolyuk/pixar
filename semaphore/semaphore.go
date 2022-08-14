package semaphore

import "sync"

type Semaphore struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewSemaphore(maxConcurrentOps uint) *Semaphore {
	return &Semaphore{make(chan struct{}, maxConcurrentOps), new(sync.WaitGroup)}
}

func (s *Semaphore) Acquire() {
	s.wg.Add(1)
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}

func (s *Semaphore) Limit() int {
	return cap(s.c)
}

func (s *Semaphore) Count() int {
	return len(s.c)
}
