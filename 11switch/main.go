package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch cases in golang")

	var randomNumber int = (rand.Int() % 6) + 1
	fmt.Println("Random number generated", randomNumber)

	switch randomNumber {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Move 2 spaces")
	case 3:
		fmt.Println("Move 3 spaces")
		//fallthrough use this to execute next case
	case 4:
		fmt.Println("Move 4 spaces")
	case 5:
		fmt.Println("Move 5 spaces")
	case 6:
		fmt.Println("Move 6 spaces")
	default:
		fmt.Print("Whoops! Not a valid move")
	}
}
