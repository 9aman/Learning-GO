package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("new custom error")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", err) // *errors.errorString is the type. See the code for errors.New() to know how this is of type error. Disclaimer : It's because a pointer to this implements the interface error.

	fmt.Println(err.Error())
}
