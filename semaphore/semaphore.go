package semaphore

import "sync"

type Semaphore struct {
	value   int
	mutex   *sync.Mutex
	waiters chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		value:   size,
		mutex:   &sync.Mutex{},
		waiters: make(chan struct{}),
	}
}

func (s *Semaphore) Wait() {
	s.mutex.Lock()

	s.value--
	if s.value >= 0 {
		s.mutex.Unlock()
		return
	}

	s.mutex.Unlock()
	<-s.waiters
}

func (s *Semaphore) Signal() {
	s.mutex.Lock()

	if s.value < 0 {
		s.waiters <- struct{}{}
	}
	s.value++

	s.mutex.Unlock()
}
