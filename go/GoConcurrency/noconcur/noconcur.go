package main

import (
	"fmt"
	"time"
	"noconcur/mypackage"
)

func main() {
	//import package
	mypackage.Print()
	//this script will keep running the first func permanently
	calc("rides")
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
