package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	// EXAMPLE SERVER
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	n, err := fmt.Fprintf(res, "Hello World")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))

	// 	// res.Write([]byte("Hello World"))
	// })

	// http.ListenAndServe(":8080", nil)

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf(fmt.Sprintf("Server running on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}

func addValues(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil // esto es para simular que podriamos devolver un error pero por ahora lo dejo null
}

func divideValues(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return 0, errors.New("Cannot divide by zero") // retornamos 0 y el error porque es lo que debermos retornar en la funcion
	}
	result = x / y
	return result, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Cannot divide by 0"))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("Sum of %d and %d is %d", 2, 2, sum))
}
