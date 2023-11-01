package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

// class Dog implements Animal
type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "Woof -> " + d.Name
}
func (d *Dog) NumberOfLegs() int {
	return 4
}

// class Gorilla implements Animal
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d *Gorilla) Says() string {
	return "Gggggrrrr -> " + d.Name
}
func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func printInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "German Sheperd",
	}

	printInfo(&dog) // para poder meter al perro como un animal debe satisfacer la interfaz de animal

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}
	printInfo(&gorilla)

}
