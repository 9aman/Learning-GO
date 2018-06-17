package main

import (
	"github.com/GoesToEleven/sampletemp/011-BET-Examples/dog"
	"fmt"
)

type canine struct {
	name string
	age int
}

func main(){
	fido := canine{
		name : "Fido",
		age : dog.YearsConv(10),
	}

	fmt.Println(fido)
	fmt.Println(dog.YearsConv1(10))
}