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

const totalRoutines = 10

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex, semaphore chan struct{}) bool {
	mutex.Lock()
	*count++
	last := *count == totalRoutines
	mutex.Unlock()

	fmt.Println("rendezvous ", goNum)
	if last {
		close(semaphore)
	} else {
		<-semaphore
	}
	fmt.Println("critical section ", goNum)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	sem := make(chan struct{})
	count := 0
	wg.Add(totalRoutines)
	//ctx := context.TODO()
	//var theLock sync.Mutex
	//sem := semaphore.NewWeighted(int64(totalRoutines))
	//theLock.Lock()
	//sem.Acquire(ctx, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &count, &wg, &m, sem)
	}
	//sem.Release(1)
	//theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
