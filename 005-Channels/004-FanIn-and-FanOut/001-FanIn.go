package main

import (
	"fmt"
	"sync"
)

func main() {
	o := make(chan int)
	e := make(chan int)

	fanin := make(chan int)
	// send
	go send1(o, e)

	// receive
	go receive1(o, e, fanin)

	fmt.Println("The received values are")
	for v := range fanin {
		fmt.Println(v)
	}

}

func send1(o, e chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(o)
	close(e)
}

func receive1(o, e <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin) // otherwise we will have deadlock as we are reading values from it in main
}

