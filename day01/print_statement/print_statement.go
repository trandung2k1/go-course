package printstatement

import "fmt"

func Init() {
	// interger %d
	// float %g, %f
	// string %s
	// bool %t
	// type %T
	var mark = 8.5
	fmt.Printf("%g\n", mark) // %f = 8.500000, %g = 8.5

	// Using print() or println() debugging

	var name string
	// fmt.Print("\nEnter your name: ")
	// fmt.Scan(&name)
	// fmt.Printf("Name: %s\n", name)

	var age int
	// fmt.Print("\nEnter your name and age: ")
	// fmt.Scanln(&name, &age)
	// fmt.Printf("Name: %s\nAge: %d", name, age)

	fmt.Println("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s\nAge: %d", name, age)
}
