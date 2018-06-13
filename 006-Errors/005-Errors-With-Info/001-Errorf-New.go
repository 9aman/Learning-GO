package main

import (
	"fmt"
	"math"
)

func main() {
	v, err := sqrt(-10)
	if err != nil {
		fmt.Printf("%T\n is the type of the error", err)
		fmt.Println("Error occurred :", err)
		return
	}
	fmt.Println("Square root is", v)
}

func sqrt(n float64) (float64, error){
	if n < 0 {
		return 0, fmt.Errorf("Value less than zero") // We can use either of these
		//return 0, errors.New("value less than zero")
	}else{
		return math.Sqrt(n), nil
	}
}