package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){

		fmt.Println("\n I will get blocked \n")
		c <- 42
		for i := 0;i < 100; i++ {
		fmt.Println("I am unblocked", i)}
		fmt.Println("unblocked first one, working on second")
		c <- 43
	}()
	// Now as soon as this func puts value in the channel it gets blocked until another routine works as a receiver and takes the value.

	fmt.Println("going to unblock anonymous func")
	fmt.Println("\n Waiting for c to be put in \n")

	// The above two values are printed out and for printing the next statement we need value in channel c, which has not been put by the anonymous func. Thus the receiver has to wait till some other go routine puts the value in the channel which can be read by this receiver.

	fmt.Println("Unblocked anonymous func",<-c)

	// now again the receiver has to wait.

	fmt.Println("received again", <-c)
	fmt.Println("Work Complete")
}
