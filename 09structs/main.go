package main

import "fmt"

func main() {
	structs()
}

// since golang does not have classes, structs are very useful in golang
// there is no inheritance in golang. i.e -> no super(), no parent
func structs() {
	var user1 User = User{"test1", "test@gmail.com", 20}

	fmt.Println("User", user1)

	user1_name := user1.Name

	fmt.Println("user 1 name : ", user1_name)

	// print key value of a custom struct, %+v enables us to print key-value

	fmt.Printf("Details of user1 are: %+v\n", user1)
}

// how to define struct

type User struct {
	Name  string
	Email string
	Age   int
}
