package main

import "fmt"

func main() {
	c := make(chan int)
	go makeChan(c)

	for i:= 0; i< 25; i++ {
		fmt.Println(<-c)
	}
}

func makeChan(c chan int) {

	count := 0

	for i :=1 ;i <=5 ; i++ { // Launching 5 channels
		count++
		go func(counter int){
			for j := 1; j <= 5; j++{
				//fmt.Println("for the counter and j", counter, j, "putting value", counter*j)
				c <- counter*j
			}
		}(count)
	}
}