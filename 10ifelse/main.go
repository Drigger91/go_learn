package main

import (
	"fmt"
	"math/rand"
)

// lexer removes the if bracket
// you cannot transfer the curly brace in the next line
func main() {
	var x int = getRandomInteger()
	if x > 3 {
		fmt.Println("greater than 3")
	} else if x == 3 {
		fmt.Println("Equals 3")
	} else {
		fmt.Println("less than 3")
	}

	// you can assign a value and check it on the flow, this can help in api requests and stuff
	if num := getRandomInteger(); num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Greater than or equal to 10")
	}

	// returnValue, err := someFunctionCall()
	// if err != nil {
	// 	handleError()
	// }
}

func getRandomInteger() int {
	return (rand.Int() / 1000) % 11
}
