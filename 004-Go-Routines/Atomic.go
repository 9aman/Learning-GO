package main

import (
	"sync/atomic"
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	wg.Add(100)
	var incrementer int64 = 0
	for i := 0; i < 100; i++ {
		go func(){
			atomic.AddInt64(&incrementer,1)
			fmt.Println("incrementer", atomic.LoadInt64(&incrementer))
			wg.Done()
			// fmt.Println(incrementer) if we read the values like this we will get data race while, LoadInt64 and AddInt64 ensures concurrency i.e. when one is reading that memory location no other goroutine can.
		}()
	}
	wg.Wait()
	fmt.Println("exited", incrementer)
}