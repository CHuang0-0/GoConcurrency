package main

import (
	"fmt"
	"time"
)

func main() {
	//make a channel
	c := make(chan string)
	//pass the channel into func
	go calc("rides", c)
	// receive a msg from the channel
	// sending and receiving block operations
	// to synchronize goroutines
	fmt.Println("use channels for goroutines to communicate with each other.\nreceiving end will wait for the sending end.\ngoroutines are synchronized.")
	for msg := range c {
		// msg, open := <-c
		// if !open {
		// 	//if channel is closed, break out the for loop
		// 	break
		// }
		fmt.Println(msg)
	}

}

func calc(myvar string, c chan string) {
	i := 0
	//create an infinite loop
	for i < 3 {
		i++
		//send a value over the channel
		c <- myvar
		// fmt.Println(i, myvar)
		time.Sleep(time.Second * 1)
	}
	//close the channel
	close(c)
}
