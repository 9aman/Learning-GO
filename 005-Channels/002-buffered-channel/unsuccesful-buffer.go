package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 42
	c <- 43

	// the buffer capacity is one but we have put two values. Thus, one of the values need to be read immediately in order for the code execution to resume.

	fmt.Println(<-c)
	// fatal error: all goroutines are asleep - deadlock! is receiver when this code is run.
	// They follow FIFO principle
}
