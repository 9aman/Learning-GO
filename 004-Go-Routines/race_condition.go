package main

import (
	"sync"
	"runtime"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(100)
	incrementer := 0
	for i := 0; i < 100; i++ {
		go func(){
			defer wg.Done()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("incrementer", incrementer)
			fmt.Println("NumGoroutine", runtime.NumGoroutine())
		}()

	}
	wg.Wait()
	fmt.Println("values of incerementer",incrementer)
	fmt.Println("No of go routines", runtime.NumGoroutine())
}
