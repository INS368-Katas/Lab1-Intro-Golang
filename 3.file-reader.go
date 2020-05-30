package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	var filename string = os.Args[1]

	var byteData, err = ioutil.ReadFile(filename)
	if err != nil {
		// Works similar to fmt.Println except it also exits the app
		log.Println("Error: Archivo no encuentrado")
	}

	// Converts byte data to string
	var data string = string(byteData)

	fmt.Println(data)
}
