package main

import (
	"fmt"
	"net/http"

	"github.com/yasniel1408/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Server running on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
