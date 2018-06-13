package main

import "fmt"

type customErr struct{
	info string
}

func (c customErr) Error() string{
	return fmt.Sprintf("Error occured, more details : %v", c.info)
}
func main() {
	err := customErr{
		info : "Name of the game",
	}
	fmt.Printf("Passsing %v error in foo.\n",err.info) // works here considering it of type customErr
	foo(err)
}

func foo(err error) {
	fmt.Println(err) // This will automatically call Error()
	//fmt.Println(err.s) // doesn't work here as .info is not attached to error. It's attached to type customErr and not error

	// What we can do here is use assertion which we used earlier to assert that the value err is also of type customErr and then use .info

	fmt.Println("Asserting the type for the error",err.(customErr).info)

	// assertion is used when multiple types are implementing an interface while conversion is converting a type into another.
}
