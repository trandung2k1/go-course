package scope

import "fmt"

// global variable
var X int = 100

// local variable
var x = 10

// Global function
func Init() {
	fmt.Println("X = ", X)

	fmt.Printf("Local: %d \n", x)
}

// local function
func init() {
	fmt.Printf("Local: %d \n", x)
}
