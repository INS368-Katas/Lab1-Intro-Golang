package main

import (
	"fmt"
	"os"
)

func main() {
	var nameArg string = os.Args[1]
	fmt.Println("Hello,", nameArg)
}