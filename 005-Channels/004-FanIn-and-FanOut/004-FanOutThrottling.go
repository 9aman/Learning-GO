package main

import (
	"sync"
	"fmt"
	"math"
	"time"
)

func main() {
	c  := make(chan int)
	c1 := make(chan int)

	go populate1(c)

	go Process(c, c1)

	for v := range c1 {
		fmt.Println(v)
	}
}

func populate1(c chan int){
	for i := 0; i<10; i++{
		c <- i
	}
	close(c)
}

func Process(c, c1 chan int){
	throughput := 10

	var wg sync.WaitGroup
	wg.Add(throughput)

	for i := 0; i< throughput; i++{ // for launching 10 go routines to read from c
		go func(){
			for v := range c{
				c1 <- operation(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c1)
}

func operation(n int) int{
	time.Sleep(time.Second	)
	return int(math.Sqrt(float64(n)))
}