package main

import "fmt"

func main() {
	c := fanIn(fillNum(40), fillNum(0))

	for i := 0; i < 10; i++ {
		fmt.Println("No received", <-c)
	}
}

func fanIn(c1, c2 <-chan int) <-chan int {

	c := make(chan int)
	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}

func fillNum(k int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; ; i++ {
			c <- k + i
		}
	}()
	return c
}

