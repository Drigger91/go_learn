package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi")
	callInputRelatedCalls()
}
func callInputRelatedCalls() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter numbers to print")

	// comma ok syntax | error ok syntax
	// you can use _ if you don't care what's return in the method other than input. For error handling you can store
	// it in a variable err

	input, _ := reader.ReadString('\n')

	fmt.Printf("User input %s\n", input)

	// conversion related code

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strconv.ParseFloat: parsing "ter\n": invalid syntax, \n will also come if you read from stdin hence using package strings

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("NUM rating, ", numRating)

}
