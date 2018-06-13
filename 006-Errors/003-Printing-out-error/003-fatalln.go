package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	// fatal does os.Exit(1) i.e. error while os.Exit(0) means success. Any non-zero value will give us error condition. All the deferred functions are not executed

	defer foo()

	f, err := os.Open("non-existing-file.txt")
	if err != nil{
		log.Fatalln("Fatal, os.Exit(1) will be executed : ", err)
	}
	defer f.Close()
}

func foo(){
	fmt.Println("WONT get printed as we enconter os.Exit(1)")
}
