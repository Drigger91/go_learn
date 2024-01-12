package main

import "fmt"

const LoginToken string = "sometoken" // capital first letter is to define this is public

func main() {
	fmt.Println("Variable file bitch")
	variableIntroduction()
	fmt.Println(LoginToken)
}
func variableIntroduction() {
	var name string = "Piyush"
	fmt.Println(name)

	var isSmallThanOne bool = 1 < 0
	// this will not work with printLn it'll take this as a common string only;
	fmt.Printf("Class of variable isSmallThanOne is %T\n", isSmallThanOne)

	var smallInteger uint8 = 255 // uint8 is 0-2^8-1 (255)
	fmt.Printf("Class and value of the variable is : %d %T\n", smallInteger, smallInteger)

	var normalInteger int = 897
	fmt.Printf("Normal Integer %d \n", normalInteger)

	var unInitialisedString string
	fmt.Printf("Value of uninitialised string is %s\n", unInitialisedString)

	// implicit type

	var unDefinedString = "Not Defined type of this"
	fmt.Println(unDefinedString)

	// no var

	x := 309 // you are only allowed to use walrus operator inside methods
	fmt.Println(x)
}
