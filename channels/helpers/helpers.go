package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {

	// generar un número aleatorio
	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(n)

	return value
}
