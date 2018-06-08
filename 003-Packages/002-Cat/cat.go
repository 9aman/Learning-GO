package cat

import "fmt"

func EAT(){
	fmt.Println("NON")
}

func RUN(){ // upper case so can be accesses from outside the directory
	fmt.Println("SLOW")
}

func cantUSEIT(){
	fmt.Println("LOCAL TO CAT")
}


func CanUSEIT(){
	fmt.Println("can be used")
}