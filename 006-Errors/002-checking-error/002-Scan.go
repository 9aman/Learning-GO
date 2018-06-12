package main

import "fmt"

func main() {
	var x, y, z string
	fmt.Println("Enter values")
	f, err := fmt.Scan(&x, &y, &z)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
	f, err = fmt.Scanf("%v %v %v",&x, &y) // if input doesn't match the format it will give an error
	if err != nil {
		fmt.Println("error encountered : ", err)
	}
	fmt.Println(f)
}