package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "Example content for reading/writing in files in go lang"

	file, err := os.Create("example.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length: ", length)

	defer file.Close()
	readFile("./example.txt")

	// for error

	readFile("random.txtx")
}

func readFile(fileName string) {
	databytes, err := os.ReadFile(fileName)

	checkNilError(err)
	fmt.Println("Data: ", string(databytes)) // same as new String(byte[]) in java

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
