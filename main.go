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
}

func saySomething() string {

	return "Something"
}

func saySomething2() (string, string) {

	return "I Say", "Something"
}
