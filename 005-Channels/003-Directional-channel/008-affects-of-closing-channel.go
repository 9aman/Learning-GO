package main

import "fmt"

func main() {

	c := make(chan int)

	go func(){
		c <- 43
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

	// If we don't close the channel then the second receive statement will create a deadlock as their is no sender which puts value into the channel. Thus, it will wait indefinitely for some goroutine to put value in channel.
	// fatal error: all goroutines are asleep - deadlock!

}
