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
