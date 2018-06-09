package main

import "fmt"

func main() {
	e := make(chan int)
	o := make(chan int)
	quit := make(chan int)

	// send
	go send(e, o, quit)

	// receive
	rec(e, o, quit)

	fmt.Println("Exited")
}

func rec(e, o, q <-chan int){
	// assignment statement is evaluated when their is value in the channel and hence then only will the case be selected.
	for {
		select {
		case v := <-e:
			fmt.Println("Even :", v)
		case v := <-o:
			fmt.Println("Odd : ", v)
		case v := <-q:
			fmt.Println("last one", v)
			return
		}
	}
}

func send(e, o, q chan<- int){
	for i := 1;i <= 10; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}