package main

import "fmt"

func main() {
	o := make(chan int)
	e := make(chan int)
	quit := make(chan int) // to stop the infinite for loop

	// send
	go sendChan(o, e, quit)

	// receive
	rece(o, e, quit)
	fmt.Println("Program exited")
}

func rece(o, e, q <-chan int){
	for {
		select {
		case v := <-o :
			fmt.Println("Odd ", v)
		case v := <-e :
			fmt.Println("Even", v)
		case v, ok := <-q :
			if !ok {
				fmt.Println("Exited, all channels are empty", v)
				return
			}else{
				fmt.Println()
			}
		}
	}
}

func sendChan(o, e, q chan<- int){
	for i := 0; i < 10; i ++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
