package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("non-existing.txt")

	if err != nil {
		fmt.Println("trying to open which doesnt exist : ", err)
	}
	fmt.Printf("%T\n", f)

	// So this simply prints to the standard out.
}
