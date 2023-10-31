package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// Map (string, interface{}) esto es para que almacene culaquier cosa pero no es recomendado

	// Map (string, string)
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other"] = "Other"

	myMap["dog"] = "Replace name"

	log.Println("The dog name is: ", myMap["dog"])
	log.Println("The other name is: ", myMap["other"])

	// Map (string, int)
	ageByName := make(map[string]int)
	ageByName["Juan"] = 30
	ageByName["Pedro"] = 40

	log.Println("The age of Juan is: ", ageByName["Juan"])
	log.Println("The age of Pedro is: ", ageByName["Pedro"])

	// Map (string, User)
	userByName := make(map[string]User)
	userByName["Juan"] = User{
		FirstName: "Juan",
		LastName:  "Perez",
	}
	log.Println("The user is: ", userByName["Juan"])

}
