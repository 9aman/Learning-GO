package main

import "fmt"

func main() {

	c := make(chan int)

	go func(){
		for i := 0; i < 5; i++{
			fmt.Println("value being entered",i)
			c <- i
		}
		close(c)
	}()

	for v := range c { // This reads from a channel till it is closed. So it's necessary for us to close the channel otherwise we will have deadlock as it will wait for the sender to put some value in the channel.
		fmt.Println("Value being read", v)
	}
	fmt.Println("Done")
}
