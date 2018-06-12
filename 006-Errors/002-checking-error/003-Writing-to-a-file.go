package main

import (
	"os"
	"fmt"
	"strings"
	"io"
)

func main() {
	f, err := os.Create("names.txt") // f is a pointer to file which we know is also of type writer as well as reader.

	if err != nil {
		fmt.Println("error creating file")
		return
	}

	defer f.Close() // where you open try to close it then and there.

	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// This gives us a string reader of type *Reader

	io.Copy(f, r) // takes writer and reader as input and reads from reader(src) and writes to the destination pointed by f. It writes till the EOF or an error is encountered.
}