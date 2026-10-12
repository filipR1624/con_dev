//reusableBarrier.go
//Copyright (C) 2026 Mr. Filip Raguz

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Filip Raguz
// Created on 05/Oct/2026
// Modified by: -
// Issues: -
//--------------------------------------------

package main

import (
	"con_dev/semaphore"
	"fmt"
	"sync"
)

// Global variable for simplicty. Describes how many threads are ran during the implementation.
const threadCount = 3

func reusableBarrier(i int, wg *sync.WaitGroup, m *sync.Mutex, count *int, s1, s2 *semaphore.Semaphore) {
	reuses := 2
	// ran in a loop to simulate the reusability
	// reuses controls HOW MANY reuses occur
	for r := range reuses {
		fmt.Println("rendezvous", i, "round", r+1)
		m.Lock()
		*count++
		last := *count == threadCount
		m.Unlock()

		if last {
			s2.Wait()
			// s2 was initialized to 1,
			// so on the first pass we are LOCKING IT here
			// (aka making sure when others reach it they block)
			s1.Signal()
		}
		s1.Wait()
		s1.Signal()
		// after all threads pass this point
		// s1 is left at 1 (aka s1 is left unlocked),
		// while s2 is at 0

		fmt.Println("critical section", i, "round", r+1)

		m.Lock()
		*count--
		theLast := *count == 0
		m.Unlock()

		if theLast {
			s1.Wait()
			s2.Signal()
		}
		// threads block here (since s2 was locked on the initial pass),
		// but after all pass, its left at 1 again,
		// thats why we need to lock s1 in theLast block
		// (aka put it back to 0)
		s2.Wait()
		s2.Signal()

		if r == 1 {
			fmt.Println(i, "finished")
		}
	}
	wg.Done()
}

func main() {
	// allows all threads to finish, intead of abruptly closing at end of main
	var wg sync.WaitGroup
	wg.Add(threadCount)
	// shared lock
	mutex := sync.Mutex{}

	// synchronization mechanisms
	sem1 := semaphore.NewSemaphore(0)
	sem2 := semaphore.NewSemaphore(1)
	count := 0

	for i := range threadCount {
		go reusableBarrier(i, &wg, &mutex, &count, sem1, sem2)
	}
	wg.Wait() //wait here until everyone is done
}
