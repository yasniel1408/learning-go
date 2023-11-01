package main

import (
	"log"
	"sort"
)

type Student struct {
	FirstName string
	LastName  string
}

func main() {

	// Array string
	var myArray []string

	myArray = append(myArray, "Juan")
	myArray = append(myArray, "Pedro")
	myArray = append(myArray, "Maria")

	log.Println("The myArray is: ", myArray)

	// Array string inicializado
	myArray2 := []string{"Juan", "Pedro", "Maria"} // dentro de los corchetes va el tamaño del array si queremos que sea de un tamaño fijo
	log.Println("The myArray2 is: ", myArray2)

	// Array Student
	var myArrayStudent []Student
	myArrayStudent = append(myArrayStudent, Student{FirstName: "Juan", LastName: "Perez"})
	myArrayStudent = append(myArrayStudent, Student{FirstName: "Pedro", LastName: "Gomez"})
	myArrayStudent = append(myArrayStudent, Student{FirstName: "Maria", LastName: "Gonzalez"})

	log.Println("The myArrayUser is: ", myArrayStudent)

	// Sort Array numbers
	var myArray3 []int
	myArray3 = append(myArray3, 1)
	myArray3 = append(myArray3, 3)
	myArray3 = append(myArray3, 2)

	log.Println("The myArray3 is: ", myArray3)
	sort.Ints(myArray3)
	log.Println("The myArray3 is: ", myArray3)

	// Print partial Array numbers
	myArray4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println("The 0-5 is: ", myArray4[0:5])
	log.Println("The 6-9 is: ", myArray4[6:9])

}
