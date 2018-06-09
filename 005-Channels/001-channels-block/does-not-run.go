package main

import "fmt"

func main() {
	c := make(chan int) // make channel of type int

	c <- 42 // Putting 42 in

	fmt.Println(<-c) // This unloads.
	// This won't work as there is no goroutine to read it. Channels receive and transmit. If their is no receiver it causes deadlock.
	// fatal error: all goroutines are asleep - deadlock!
	// In this example we don't have any other goroutine which works as the receiver.
}
