package main

import (
	"fmt"
)

func main()  {

	// bi-directional channel
	d := make(chan int)
	// this makes chan d into a send only channel. we can only recive from this channel
	dr := make(<- chan int)
	// this makes chan d into a recieve only channel. we can only send to the channel
	ds := make(chan <- int)
	
 
	fmt.Println("-----")
	fmt.Printf("d\t%T\n", d)
	fmt.Printf("dr\t%T\n", dr)
	fmt.Printf("ds\t%T\n", ds)

}