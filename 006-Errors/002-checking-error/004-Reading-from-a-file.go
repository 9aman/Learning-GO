package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error openeing file", err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f) // takes reader as input and returns a byte stream and error
	if err != nil {
		fmt.Println("file can't be read", err)
	}
	fmt.Println(string(bs))
}