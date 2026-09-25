//rendezvous.go Template Code
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
// Author: Joseph Kehoe (joseph.kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Filip Raguz
// Help: Thomas Radulescu
// Issues: -
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// duration is the maximum amount of seconds a thread can sleep
const duration = 5

type semaphore struct {
	n chan struct{}
}

// initialize creates a buffered channel of capacity 1 to mimic a semaphore
func initialize() semaphore {
	return semaphore{
		n: make(chan struct{}, 1),
	}
}

// signal mimics incrementing a semaphore
func (s semaphore) signal() {
	s.n <- struct{}{}
}

// wait mimics decrementing a semaphore
func (s semaphore) wait() {
	<-s.n
}

// rendezvous requires all threads to finish PartA before moving on to PartB, aka they need to "rendezvous"
func rendezvous(thread int, wg *sync.WaitGroup, semA, semB *semaphore) {
	// wait a random amount of seconds
	n := time.Duration(rand.IntN(duration))
	time.Sleep(n * time.Second)

	fmt.Println("PartA: ", thread)
	semA.signal() // unblocks the other threads
	semB.wait()   // blocks/waits for the other thread
	fmt.Println("PartB: ", thread)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	s1 := initialize()
	s2 := initialize()

	for i := range 2 {
		wg.Add(1)
		if i == 0 {
			go rendezvous(i, &wg, &s1, &s2)
		} else {
			go rendezvous(i, &wg, &s2, &s1)
		}

	}

	wg.Wait()
}
