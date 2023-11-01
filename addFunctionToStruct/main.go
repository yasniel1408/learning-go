package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string // esta propiedad es publica pero si es con minuscula es privada, igual la podemos setear pero si decimos p.name = "Juan" es como si fuera el set
}

// de esta manera metemos la funcion dentro de la estructura para cuando se use la estructura se pueda usar la funcion
func (p *Person) printName() string {
	log.Println("The name is: ", p.Name)
	return p.Name
}

func main() {

	var p Person
	p.Name = "Juan"

	p2 := Person{
		Name: "Pedro",
	}

	fmt.Println("The person is: ", p.Name)
	fmt.Println("The person is: ", p2.Name)

	fmt.Println("The person name is: ", p.printName())

}
