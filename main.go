package main

import "fmt"

func main() {
	fmt.Println("Init World")

	// String
	var wathToSay string

	wathToSay = "Godbye cruel world"

	fmt.Println(wathToSay)

	// Integer
	var number int
	number = 42
	fmt.Println("El numero es: ", number)

	//function
	something := saySomething()
	fmt.Println("The function return something: ", something)

	something1, something2 := saySomething2()
	fmt.Println("The function return something: ", something1, something2)

	//Pointer
	color := "Green"
	fmt.Println("BEFORE: The color is: ", color)
	changeUsePointer(&color)
	fmt.Println("AFTER: The color is: ", color)

}

func saySomething() string {
	return "Something"
}

func saySomething2() (string, string) {
	return "I Say", "Something"
}

func changeUsePointer(e *string) {
	fmt.Println("The pointer is: ", e) // The pointer is:  0xc000010200
	fmt.Println("The value is: ", *e)  // The pointer is:  Green

	newValue := "Red"
	*e = newValue
}
