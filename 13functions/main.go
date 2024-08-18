package main

import (
	"errors"
	"fmt"
	"math"
)

// main act as a entry point for your go file
func main() {
	fmt.Println("This is intro to functions")

	fmt.Println("2 + 3 = ", add(2, 3))

	fmt.Println("2 + 3 + 5 + 7 = ", addMulti(2, 3, 5, 7))

	logErrorIfAny(3)
	logErrorIfAny(7)
}
func logErrorIfAny(a int) {
	ans, err := handleError(a)

	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println("return value", ans)
	}
}
func add(a int, b int) int {
	return (a + b)
}

func addMulti(values ...int) int {
	ans := 0
	for index := range values {
		ans += values[index]
	}
	return ans
}

func handleError(a int) (int, error) {
	if a == 3 {
		// apparantly you cannot use caps in error message
		return math.MinInt, errors.New("three not supported for this method")
	}
	return a + 5, nil
}
