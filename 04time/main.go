package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hey there from time module!")
	callTimeRelatedFunctions()
}

// this is how you write documentation in go, when you over the functionCall() this will give the description
func callTimeRelatedFunctions() {
	presentTime := time.Now()
	fmt.Println("Current time", presentTime)

	//This is some go related bs that you have to use this specific date for formatting only, examples are below.

	formatedTime := presentTime.Format("01-02-2006")
	fmt.Println("Format time", formatedTime)

	// include day and yes Wednesday is also a fixed value

	timeDetailsWithDay := presentTime.Format("01-02-2006 Wednesday")
	fmt.Println("Day details", timeDetailsWithDay)

	dateTimeDetails := presentTime.Format("01-02-2006 15:04:05")
	fmt.Println("Date time", dateTimeDetails)

	// create custom date, same operations as time.Now()

	//var myDate time.Time = time.Date(2020, time.March, 8, 0, 0, 0, 0, time.UTC)
	myDate := time.Date(2020, time.March, 8, 0, 0, 0, 0, time.UTC)
	fmt.Println("my custom date ", myDate)

}
