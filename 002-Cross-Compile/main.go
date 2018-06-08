package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %v\nARCH: %v\n", runtime.GOOS, runtime.GOARCH)
}