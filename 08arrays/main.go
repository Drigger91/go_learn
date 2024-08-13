package main

import (
	"fmt"
	"sort"
)

func main() {
	var list [4]string
	list[0] = "hi"
	list[3] = "hello"

	fmt.Println("list", list)
	fmt.Println("length of list", len(list))

	// you can initialise arrays by this syntax:

	var listNew = [3]string{"new", "old"}
	fmt.Println("listNew", listNew)

	callSlicesRelatedMethods()
}

func callSlicesRelatedMethods() {
	fmt.Println("<---------Slices----------->")
	var sliceExample = make([]string, 4)
	sliceExample[0] = "first"
	// append is push_back
	sliceExample = append(sliceExample, "example")
	for i := 0; i < len(sliceExample); i++ {
		fmt.Println(sliceExample[i])
	}

	// second syntax
	fmt.Println("<<<<<-second->>>>>>")
	var newSlice = []string{}

	newSlice = append(newSlice, "newFirst")

	// slice[include : exclude]
	for i := 0; i < len(newSlice); i++ {
		fmt.Println(newSlice[i])
	}
	fmt.Println("-<<<Third>>>-")
	sort.Strings(sliceExample)
	for i := 0; i < len(sliceExample); i++ {
		fmt.Println(sliceExample[i])
	}

	fmt.Println("Remove values from slices based on index")

	var courses = []string{"Course1", "Course2", "Course3", "Course4"}
	for i := 0; i < len(courses); i++ {
		fmt.Println(courses[i])
	}

	fmt.Println("---------------")
	// use append to remove the values
	var indexToDelete int = 2
	courses = append(courses[:indexToDelete], courses[indexToDelete+1:]...)
	for i := 0; i < len(courses); i++ {
		fmt.Println(courses[i])
	}
}
