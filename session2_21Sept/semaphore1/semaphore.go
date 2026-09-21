// Author: Filip Raguz
// Date: 21/09/2026
// Who helped me:
// - Kristian Kesar
// - Thomas Radulescu

package main

import (
	"fmt"
	"sync"
	"time"
)

type semaphore struct {
	theCounter chan struct{}
}

func initialize(n int) *semaphore {
	return &semaphore{theCounter: make(chan struct{}, n)}
}

func acquite(sem *semaphore) {
	sem.theCounter <- struct{}{}
}

func release(sem *semaphore) {
	<-sem.theCounter
}

func main() {
	maxGoroutines := 5
	semaphore := initialize(maxGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			acquite(semaphore)
			defer func() { release(semaphore) }()

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
