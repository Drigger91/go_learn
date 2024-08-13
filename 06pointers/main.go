package main

import "fmt"

func main() {
	var myNum int = 23
	var ptr = &myNum

	//pointers can be used for pass by reference

	fmt.Println("Pointer points to address", ptr)
	fmt.Println("Pointer holds the value", *ptr)

	// this is how pointers ensures that the operations performed are on the actual values rather than on the copies of values
	*ptr = *ptr * 2
	fmt.Println("Updated value", myNum)
}
