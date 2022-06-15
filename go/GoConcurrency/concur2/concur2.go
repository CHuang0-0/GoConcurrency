package main

import (
	"fmt"
	"sync" // use waitgroup to finish goroutines
	"time"
)

func main() {
	fmt.Println("use the waitgroup from package sync.")
	var wg sync.WaitGroup
	wg.Add(1) //add by 1: 1 goroutine to wait for
	fmt.Println("the goroutine will stop after count=3.")
	//this script will keep running the first func permanently
	go func() {
		calc("rides")
		wg.Done()
	}()
	//calc("eats")
	wg.Wait() //block until the wait count is 0
}

func calc(myvar string) {
	i := 0
	//while loop: count up to 3
	for i < 3 {
		i++
		fmt.Println(i, myvar)
		time.Sleep(time.Second * 1)
	}
}
