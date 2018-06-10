package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	c1 := make(chan int)

	go populate(c)

	go process(c1, c)

	fmt.Println("entered in main")

	counter := 0
	for v := range c1 {
		counter++
		fmt.Println("value of sqareroot of nos from 1 to ", "is", v)
	}

}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func process(c1, c chan int) {

	var wg sync.WaitGroup

	for v := range c {
		wg.Add(1)
		go func(n int) {
			c1 <- preprocess(n)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c1)
}

func preprocess(n int) int {
	time.Sleep(time.Second)
	return int(math.Sqrt(float64(n)))
}

