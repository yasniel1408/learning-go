package main

import (
	"fmt"
	"log"
	"time"
)

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

	// Struct
	type User struct {
		ID        int
		FirstName string
		LastName  string
		address   string // Si la variable es minuscula es privada y si con mayuscula es publica, lo mismo para las funciones en un archivo
		Age       int
		BirthDate time.Time
	}
	user := User{
		ID:        1,
		FirstName: "Juan",
		LastName:  "Perez",
		// Address:  "Calle 123",
		Age:       30,
		BirthDate: time.Now(),
	}
	log.Println("The user is: ", user)
	log.Println("The user name is: ", user.FirstName)
	log.Println("The address is: ", user.address)

	// Maps
	type Person struct {
		Name string // esta propiedad es publica pero si es con minuscula es privada, igual la podemos setear pero si decimos p.name = "Juan" es como si fuera el set
	}
	var p Person
	p.Name = "Juan"

	p2 := Person{
		Name: "Pedro",
	}

	fmt.Println("The person is: ", p.Name)
	fmt.Println("The person is: ", p2.Name)

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
