package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	f1, err1 := os.Create("log.txt")
	if err1 != nil{
		log.Println("error creating file", err1)
		// if the location for log is not specified it uses stdout as the default. it also prints the time stamp.
	}
	defer f1.Close()

	log.SetOutput(f1) // Now it will use log.txt for logging.

	f2, err2 := os.Open("non-existing-file.txt") // non-existing file
	if err2 != nil {
		log.Println("error encountered", err2)
	}
	defer f2.Close()

	// logs will be present in log.txt for the second case and in stdout for the first case.
	fmt.Println("Errors logged in log.txt")
}
