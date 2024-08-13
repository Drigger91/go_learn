package main

import "fmt"

func main() {
	var nicknames map[string]string = make(map[string]string)

	nicknames["swapnil"] = "baniya"
	nicknames["zaid"] = "zaid"
	nicknames["mayank"] = "mayank"
	nicknames["piyush"] = "tiwari"
	nicknames["tobe"] = "deleted"

	fmt.Println("Map ", nicknames)

	printSpace()

	delete(nicknames, "tobe")

	fmt.Println("Map ", nicknames)

	// iterate map

	printSpace()
	// while print it gives ordered values but while iterating it gives values in order by they entered
	for key, val := range nicknames {
		fmt.Printf("%v => %v\n", key, val)
	}
}

func printSpace() string {
	fmt.Println("<----->")
	return "hey"
}
