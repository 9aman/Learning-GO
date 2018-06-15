package main

import (
	"github.com/GoesToEleven/sampletemp/007-Documentation/001-Dog"
	"fmt"
)

type canine struct{
	name string
	age int

}

func main() {
	d1 := canine{
		name : "Tommy",
		age : Dog.Year(5),
	}
	fmt.Println(d1)
}
