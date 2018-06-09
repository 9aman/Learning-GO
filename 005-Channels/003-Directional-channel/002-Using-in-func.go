package main

import "fmt"

func main() {

	c := make(chan int)

	// receive
	//fmt.Println("going for bar go routine")
	//go bar(c) This works
	// bar(c) This doesn't work

	// if we put receive before send, because of bar belonging to the same goroutine as that of the main, the main will be blocked till it receives the value from channel and hence we will never reach to go foo(c) statement, thus causing deadlock. Although if we put go bar(c) then it will work because then it will have a different goroutine than that of main which will allow normal execution of main.

	// send
	fmt.Println("going for foo go routine")
	go foo(c)

	fmt.Println("going for bar go routine")
	bar(c)
	//go bar(c)
	// The main goroutine won't wait for these goroutines to execute and the program will be exited. To avoid this either we can use WaitGroups or we can use bar(c) instead of go bar(c) as now the main has to wait for the bar to complete which in turn has to wait for foo to complete.

	fmt.Println("Code Exited, channel cleared")
}

//send
func foo(c chan<- int){
	fmt.Println("loading value in channel")
	c <- 4345
}

//receive
func bar(c <-chan int){
	fmt.Println("value received", <-c)
}
