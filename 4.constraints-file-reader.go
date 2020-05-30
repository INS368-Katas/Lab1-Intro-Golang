package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var filePath string = os.Args[1]

	var byteData, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln("Archivo: No encontrado")
	}

	// Converts byte data to string
	var data string = string(byteData)
	
	var input = strings.NewReader(data) // Converts string data to reader
	scanner := bufio.NewScanner(input) // Converts reader data to scanner

	// .Scan() advances the Scanner to the next token
	for scanner.Scan() {
		// .Text() returns the current token into string format
		var line string = scanner.Text()

		if len(line) <= 10 {
			var errorMessage string = line + " (os.Stderr)"
			fmt.Fprintln(os.Stderr, errorMessage)
		} else {
			var outputMessage string = line + " (os.Stdout)"
			fmt.Fprintln(os.Stdout, outputMessage)
		}
	}
}
