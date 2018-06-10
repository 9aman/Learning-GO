package main

import (
	"fmt"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for i := 6; i <= 15; i++ {
			c2 <- i
		}
		close(c2)
	}()
	f1 := false
	f2 := false
	for {
		if f1 && f2 {
			fmt.Println("for exited")
			break
		}
		select {
		case v, flag1 := <-c1: // flag catches whether the channel is closed or not
			if flag1 {
				fmt.Println("first case", v)
			} else {
				f1 = true
			}
		case v, flag2 := <-c2: // flag catches whether the channel is closed or not
			if flag2 {
				fmt.Println("second case", v)
			} else {
				f2 = true
			}
		}
	}
	fmt.Println("Program exited")
}
