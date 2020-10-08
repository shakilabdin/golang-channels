package main

import (
	"fmt"
)

func main()  {

	// send a go routine to recieve the value 
	c := make(chan int)
	go func() {
		c <- 42
		c <- 43
		c <- 43
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	// channel buffer - 1 says this channel can hold 1 buffer / cannot send more than 1
	// if we added more buffer FIFO
	d := make(chan int, 1)
	d <- 63
	// d <- 50 this would be blocked
	fmt.Println(<- d)

}