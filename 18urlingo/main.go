package main

import (
	"fmt"
	"net/url"
)

const validUrl string = "http://localhost:3001/app/learn?course=1&chapter=21 "

func main() {
	fmt.Println("Url is ", validUrl)

	// parse url

	result, _ := url.Parse(validUrl)

	fmt.Println("Parse result ", result.Path, result.Host, result.Query(), result.RawQuery, result.Port())
	fmt.Println("")

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost:3000",
		Path:     "/api/course",
		RawQuery: "token=TKN&source=SRC",
	}

	fmt.Println("Another url", partsOfUrl.String())
}
