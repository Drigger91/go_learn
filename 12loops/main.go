package main

import "fmt"

func main() {
	fmt.Println("<<--Loops-->>")

	var nums []int

	for i := 1; i < 16; i++ {
		nums = append(nums, i)
	}

	fmt.Println("nums: ", nums)

	for i := range nums {
		fmt.Println(nums[i])
	}

	// for each

	for index, value := range nums {
		fmt.Printf("Value at index %d is %d\n", index, value)
	}

	var index int = 1
	fmt.Println("---------------------------")
	// while loop
	// there is nothing nums[index++] in golang
	for index < 10 {
		if index == 3 {
			goto gotostatement
		}
		fmt.Println("Value : ", nums[index])
		index++
	}

	// go to statements

gotostatement:
	fmt.Println("this is a gotostatement")
}
