package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Job    string `json:"job"`
	HasDog bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"name": "Ed",
			"age": 21,
			"job": "Programmer",
			"has_dog": true
		},
		{
			"name": "John",
			"age": 21,
			"job": "Programmer",
			"has_dog": false
		},
		{
			"name": "Doe",
			"age": 21,
			"job": "Programmer",
			"has_dog": true
		}
	]
	`
	// READ JSON
	var unmarshalled []Person

	// con esta funcion se puede convertir un json a un slice de structs
	// convertimos el json a byte y le pasamos la variables donde queremos recuperar los datos
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// WRITE JSON
	var myJson2 []Person

	// creamos un slice de structs
	p1 := Person{
		Name:   "Ed",
		Age:    21,
		Job:    "Programmer",
		HasDog: true,
	}

	var p2 Person
	p2.Name = "John"
	p2.Age = 21
	p2.Job = "Programmer"
	p2.HasDog = false

	myJson2 = append(myJson2, p1)
	myJson2 = append(myJson2, p2)

	// con esta funcion se convierte un slice o array de structs a json
	newJson, err := json.MarshalIndent(myJson2, "", " ")

	if err != nil {
		log.Println("Error marshalling json", err)
	}

	log.Printf("newJson: %v", string(newJson))

}
