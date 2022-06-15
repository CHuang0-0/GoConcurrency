package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("creating a goroutine that runs concurrently.")
	fmt.Println("the two programs will run side by side.")
	//add a goroutine so the first program runs in the background
	go calc("rides")
	calc("eats")
}

func calc(myvar string) {
	i := 0
	//create an infinite loop
	for {
		i++
		fmt.Println(i, myvar)
		time.Sleep(time.Second * 1)
	}
}
