package semaphore

import "sync"

type semaphore struct {
	value   int
	mutex   *sync.Mutex
	waiters chan struct{}
}

func NewSemaphore(size int) *semaphore {
	return &semaphore{
		value:   size,
		mutex:   &sync.Mutex{},
		waiters: make(chan struct{}),
	}
}
