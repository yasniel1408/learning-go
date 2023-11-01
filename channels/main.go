package main

import (
	"log"

	"github.com/yasniel1408/example/helpers"
)

const numPool = 1000000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)

	intChan <- randomNumber
}

func main() {
	intChan := make(chan int) // crear un canal de enteros
	defer close(intChan)      // cerrar el canal

	go CalculateValue(intChan) // ejecutar la función CalculateValue en un hilo aparte

	num := <-intChan // recibir el valor del canal
	log.Println(num)
}
