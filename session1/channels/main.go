package main

import (
	"fmt"
	"time"
)

func main() {
	basicChannel()
	time.Sleep(2 * time.Second)
}

func basicChannel() {
	c := make(chan int)
	go func() {
		c <- 565656
	}()
	go func() {
		var x int
		var ok bool
		x = <-c
		fmt.Println(x)
		close(c)
		x, ok = <-c
		if !ok {
			fmt.Println("Closed")
			fmt.Println(x)
		}
	}()
}
