package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c1 := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(".....................")
	fmt.Printf("%T\n", c)

	sc := make(chan <- int, 2) // This of type send only i.e. we can only send values to it and can't receive it.

	bc := make(<-chan int) // This is receive only channel i.e. we can only receive values from it.

	fmt.Printf("%T\n",sc)
	fmt.Printf("%T\n",bc)

	// We can assign a more general channel to a specific channel while we can't assign a specific channel to a general channel i.e. we can't assign a directional(specific in functionality) channel to non-directional(general in functionality) channel while we can assign a non-directional channel to a directional channel

	fmt.Println("Value of c", c)
	fmt.Printf("Value of sc %v and its type %T\n",sc,sc)
	sc = c   // This does work
	fmt.Printf("Value of sc %v and its type %T\n",sc,sc)

	fmt.Printf("Value of bc %v and its type %T\n",bc,bc)
	bc = c   // This does work
	fmt.Printf("Value of bc %v and its type %T\n",bc,bc)

	// c1 = sc This doesn't work, going from specific to general is not allowed
	fmt.Println("Value of c1", c1)
	// c1 = bc This doesn't work, going from specific to general is not allowed


}
