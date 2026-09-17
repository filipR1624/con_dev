package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var runningTotal int
	var safeTotal atomic.Int64
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func() {
			for i := 0; i < 50000; i++ {
				runningTotal++
				safeTotal.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("runningTotal: %d\n", runningTotal)
	correctTotal := safeTotal.Load()
	fmt.Printf("correctTotal: %d\n", correctTotal)
}
