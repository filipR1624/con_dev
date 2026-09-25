// Author: Filip Raguz
// Date: 21/09/2026
// Who helped me:
// - Thomas Radulescu

package main

import (
	"fmt"
	"sync"
	"time"
)

var count, threadCount int

func initialize() chan bool {
	return make(chan bool)
}

func WorkWithRendezvous(wg *sync.WaitGroup, thread int, b chan bool, m *sync.Mutex) {
	fmt.Println("PartA:", thread)

	//Rendezvous
	m.Lock()
	count++
	last := count == threadCount // if count == threadCount, last will be true
	m.Unlock()

	// if last variable
	if last {
		close(b) // close the channel (unblock all waiting threads)
	} else {
		<-b // read from channel (block a thread)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("PartB:", thread)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	barrier := initialize() // initialize the "barrier" (channel)
	mutex := &sync.Mutex{}  // for mutual exclusion of count
	// global, for convenience
	threadCount = 5
	count = 0

	wg.Add(threadCount)
	for n := range threadCount {
		go WorkWithRendezvous(&wg, n, barrier, mutex)
	}
	wg.Wait() //wait until everyone is done

}
