package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello")
	// n is no of bytes and error is the error returned.
	// Here the value of n will be 6, one for each character i.e. 5 and one for newline character.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
