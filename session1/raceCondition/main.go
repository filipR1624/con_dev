package main

import (
	"fmt"
	"time"
)

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	race()
	fmt.Println("Next")
	time.Sleep(2000 * time.Millisecond)
	threads(5)
	time.Sleep(2000 * time.Millisecond)
}

func threads(num int) {
	for i := 0; i < num; i++ {
		go func(num int) {
			fmt.Printf("this is thread %d \n", num)
		}(i)
		fmt.Println("launched:", i)
	}
}

func race() {
	i := 0
	for _, s := range []string{"a", "b", "c", "d", "e"} {
		i = i + 1
		go func(num int) {
			fmt.Printf("this is thread %d letter %s \n", num, s)
		}(i)
		fmt.Println("launched:", i)
	}
}
