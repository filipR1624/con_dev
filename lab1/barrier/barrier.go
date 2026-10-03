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
// Modified by: Filip Raguz
// Issues: -
// Help: Thomas Radulescu
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
)

const totalRoutines = 3

type semaphore struct {
	emaphore chan struct{}
}

func (s *semaphore) initialize(size int) {
	s.emaphore = make(chan struct{}, size)
}

// barrier is an implementation of rendezvous for n number of threads.
// n-1 threads will block until the nth one unblocks them.
func barrier(i int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex, s *semaphore) {
	fmt.Println("PartA", i)
	mutex.Lock()
	*count++
	if *count == totalRoutines {
		mutex.Unlock()
		s.emaphore <- struct{}{}
		<-s.emaphore
	} else {
		mutex.Unlock()
		<-s.emaphore
		s.emaphore <- struct{}{}
	}

	fmt.Println("PartB", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	var m sync.Mutex
	var s semaphore
	s.initialize(0)
	count := 0
	for i := range totalRoutines {
		go barrier(i, &count, &wg, &m, &s)
	}
	wg.Wait() //wait for everyone to finish before exiting
}
