package main

import (
	"fmt"
	"time"
)

func main() {
	//make 2 channels
	c1 := make(chan string)
	c2 := make(chan string)
	// 2 goroutines with infinite loops
	go func() {
		for {
			c1 <- "every 1 sec"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "every 2 secs"
			time.Sleep(time.Second * 2)
		}
	}()

	// use select statement
	fmt.Println("use select statement to receive msg from whichever channel is ready.")
	// infinite for loops
	for {
		select {
		case msg1 := <-c1:
			// receive from channel 1
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
