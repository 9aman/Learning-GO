package main

import "fmt"

func fill1(c chan int, count int){
	for i := 0; i < 5 ; i++ {
		c <- i
	}
	count++
}
func main() {
	c := make(chan int)
	count := 0
	go fill1(c, count)
	go fill1(c, count)

	for i:= 0; i< 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("count value", count)
	fmt.Println("Exited")
}
