//semaphore implementation
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
// Issues: none that I know of (for now)
//--------------------------------------------

package semaphore

import "sync"

// Semaphore is a structure used to hold data related to a semaphore.
type Semaphore struct {
	value   int
	mutex   *sync.Mutex
	waiters chan struct{}
}

// NewSemaphore initializes a semaphore with the specified value. It returns a pointer to the resulting semaphore.
func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		value:   size,
		mutex:   &sync.Mutex{},
		waiters: make(chan struct{}),
	}
}

// Wait decrements the value of the semaphore. If it becomes less then 0, it blocks the thread.
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

// Signal increments the value of the semaphore. This function call is non-blocking.
// If there are any blocking (waiting) threads, one is awoken.
func (s *Semaphore) Signal() {
	s.mutex.Lock()

	if s.value < 0 {
		s.waiters <- struct{}{}
	}
	s.value++

	s.mutex.Unlock()
}
