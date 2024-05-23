package main 

import (
	"fmt"
	"time"
)


func main() {
	hello()
	date()
}

//prints welcome message
func hello() {
	fmt.Println("Welcome to the playground!")
}

//prints the date
func date() {
	fmt.Println("The time is", time.Now())
}
