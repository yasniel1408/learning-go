package main

import (
	"log"

	"github.com/yasniel1408/mypaquete/helpers"
)

func main() {

	var myVar helpers.SomeStruct
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)

}
