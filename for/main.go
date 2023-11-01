package main

import "log"

func main() {
	// For
	for i := 0; i < 5; i++ {
		log.Println("The number is: ", i)
	}

	animals := []string{"dog", "cat", "horse", "bird", "fish"}
	for i, animal := range animals {
		log.Println("The animal is: ", animal, " and the position is: ", i)
	}

	for _, animal := range animals { // ignore index i
		log.Println("The animal is: ", animal)
	}

	// For with Map......ESTO ES MUY UTIL
	animales := make(map[string]string)
	animales["dog"] = "Samson"
	animales["cat"] = "Michi"
	animales["other"] = "Other"

	for key, value := range animales {
		log.Println("The animal is: ", key, " and the name is: ", value)
	}

	// For string
	var myString string = "Hello World"
	for index, letter := range myString {
		log.Println("The letter is: ", string(letter), " and the position is: ", index)
	}

	// For in struct array
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{FirstName: "Juan", LastName: "Perez", Email: "juan@gmail.com", Age: 30})
	users = append(users, User{FirstName: "Pedro", LastName: "Gomez", Email: "pedro@gmail.com", Age: 40})
	users = append(users, User{FirstName: "Maria", LastName: "Gonzalez", Email: "maria@gmail.com", Age: 50})

	for i, user := range users {
		log.Println("The user is: ", users[i]) // is same
		log.Println("The user is: ", user)     // is same
	}

}
