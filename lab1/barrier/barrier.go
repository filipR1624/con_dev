//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

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
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Filip Raguz C00301624
// Issues: -
// Help: Thomas Radulescu
//--------------------------------------------

package main

import (
	"con_dev/semaphore"
	"fmt"
	"sync"
)

// Global for simplicity. Symbolizes total amount of threads (go Routines).
const totalRoutines = 3

// barrier is an implementation of rendezvous for n number of threads.
// n-1 threads will block until the nth one unblocks them.
func barrier(i int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex, sem *semaphore.Semaphore) {
	fmt.Println("PartA", i)

	mutex.Lock()
	*count++
	last := *count == totalRoutines
	mutex.Unlock()

	if last {
		sem.Signal()
	}

	sem.Wait()
	sem.Signal()

	fmt.Println("PartB", i)
	wg.Done()
}

// Implementation of barrier.
// WaitGroups are used to allow threads to finish their work, otherwise the program will return abruptly.
func main() {
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	var m sync.Mutex
	s := semaphore.NewSemaphore(0)
	count := 0
	for i := range totalRoutines {
		go barrier(i, &count, &wg, &m, s)
	}
	wg.Wait() //wait for everyone to finish before exiting
}
