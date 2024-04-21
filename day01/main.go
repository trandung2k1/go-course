package main

import (
	"day01/helloworld"
	printstatement "day01/print_statement"
	"day01/scope"
	"day01/variables"
	"fmt"
)

func main() {
	helloworld.PrintMessage()
	variables.Init()

	// Function global
	scope.Init()

	scope.X += 1
	// Variable global
	fmt.Println(scope.X)

	// Not accessible
	// scope.init()

	printstatement.Init()

}
