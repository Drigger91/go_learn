package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello from memory management!")
	callRuntimeRelatedFunction()
}

func callRuntimeRelatedFunction() {
	cpuUsage := runtime.NumCPU()
	fmt.Println("Number of CPU ", cpuUsage)
}
