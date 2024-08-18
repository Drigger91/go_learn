package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	deferFunction()
	workWithoutDefer()
}

// defer keyword:
// defer essentially puts the task at the last line of the execution context of the function. It acts as LIFO.
func deferFunction() error {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error occurred")
		return err
	}
	defer file.Close()

	if someCondition() {
		fmt.Println("Error because of some condition, returning")
	}

	// further file processing

	return nil
}

func someCondition() bool {
	return true
}

// if the same function was to be executed w/o defer it would be like:
// here we have to ensure manually that the file we opened is closed every time we are exiting the function context

func workWithoutDefer() error {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error occurred")
		return err
	}

	if someCondition() {
		file.Close()
		fmt.Println("Error because of some condition, returning")
		return errors.New("error occurred")
	}

	// further file processing

	file.Close()
	return nil
}
