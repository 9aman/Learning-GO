package main

import "fmt"

type person struct{
	name string
	age string
}

type human interface {
	speak()
}

func (p *person) speak(){
	fmt.Println("Hi, my name is",p.name,"and my age is",p.age) // We can directly use p.age and p.name as p is of type struct which allows both p.age and (*p).age
}

func saySomething(h human){
	h.speak()
}

func main(){
	p := person{
		name: "James Bond",
		age: "32",
	}
	saySomething(&p)
	//saySomething(p)// this doesn't work
}
