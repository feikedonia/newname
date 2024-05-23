package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number()
}

//prints a random number
func number() {
	fmt.Println("My favorite number is: ", 
	rand.Intn(10))
}

