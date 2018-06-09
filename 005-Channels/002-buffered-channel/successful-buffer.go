package main

import "fmt"

func main() {
	c := make(chan int, 1) // 1 is the buffer size
	// this means that the buffer can hold one value without it being immediately received by a receiver. Thus after putting in the value, the sender execution is not blocked by the channel, it resumes it's execution and allows the receiver to receive at it's own pace.

	c <- 42
	// Now since the channel has a buffer the sender and receiver and not locked and can move on with their execution fetching value from the channel when it pleases them (If the channel buffer gets filled it will act like an unbuffered channel like we saw earlier.)
	fmt.Println(<-c)
}
