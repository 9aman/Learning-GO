package main

import (
	"sync"
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main(){
	fmt.Printf("CPU %v\n", runtime.NumCPU())
	wg.Add(2)
	go foo()
	fmt.Println("No of go routines:", runtime.NumGoroutine())
	go bar()
	fmt.Println("No of go routines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("No of go routines:", runtime.NumGoroutine())
	fmt.Println("The End")
}

func foo(){
	defer wg.Done()
	fmt.Println("Entered in foo")
}
func bar(){
	defer wg.Done()
	fmt.Println("Entered in bar")
}