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

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")
		//wait here
		barrier <- true
		fmt.Println("StuffOne - PartB")
		wg.Done()
		return true
	}
	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")
		//wait here
		<-barrier
		fmt.Println("StuffTwo - PartB")
		wg.Done()
		return true
	}
	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()
	wg.Wait() //wait here until everyone (2 go routines) is done

}
