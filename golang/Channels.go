package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func deadLock() {
	var c = make(chan int)
	c <- 1
	var i = <-c // This operation will block the main goroutine waiting for another goroutine to write to the channel
	fmt.Println(i)
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
