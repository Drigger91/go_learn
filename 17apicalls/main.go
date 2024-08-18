package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/users"

const inValid = "invalidUrlToCheckerror"

func main() {
	fmt.Println("Will call url", url)
	callApi()
	callApiInvalid()
}

func callApiInvalid() {
	response, err := http.Get(inValid)
	if response != nil {
		defer closeServer(*response)
	}

	checkNilError(err)

}

func callApi() {
	response, err := http.Get(url)

	checkNilError(err)
	defer closeServer(*response)

	fmt.Printf("Response is of type %T\n", response)

	responseBody, err := io.ReadAll(response.Body)

	checkNilError(err)

	fmt.Println("Response from API is: ", string(responseBody))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func closeServer(res http.Response) {
	if res.Body == nil {
		fmt.Println("Response is nil")
		return
	}
	fmt.Println("Closing server request by calling response.Body.Close()")

	res.Body.Close()
}
