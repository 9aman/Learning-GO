package main

import (
	"sync"
	"runtime"
	"fmt"
)

func main(){
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(100)
	incrementer := 0

	for i := 0; i < 100; i++ {
		go func(){
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			// mu.Unlock() if we put our unlock here we will get data race as in the next line we are accessing the value of incrementer in the println which is reading or writing in a variable by multiple goroutines at the same time.
			fmt.Println("incrementer", incrementer)
			mu.Unlock()
			fmt.Println("NumGoroutine", runtime.NumGoroutine())
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("values of incerementer",incrementer)
	fmt.Println("No of go routines", runtime.NumGoroutine())
}
