package main

import (
	"os"
	"log"
)

func main() {
	f, err := os.Open("non-existing.txt")

	if err != nil {
		log.Panicln("\nPanick !!!!!!!!!!,", err)
	}
	f.Close()
}
