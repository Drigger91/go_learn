package main

import "fmt"

func main() {

}

type User struct {
	PublicField  int
	privateField int
}

// this user is passed as a copy, for passing by reference use ptrs
func (user User) GetPublicField() {
	fmt.Println("Public field : ", user.PublicField)
}

// you can use the methods starting with capital letters in other packages. This will help to set/get private fields in the struct
func (user User) GetPrivateField() {
	fmt.Println("Private field : ", user.privateField)
}
