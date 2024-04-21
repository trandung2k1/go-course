package variables

import (
	"fmt"
	"math"
)

func Init() {
	var i int // default value = 0
	fmt.Println(i)

	var num = 10
	// change value
	num = 20
	fmt.Printf("%d\n", num)

	// shorthand
	class := "12B3"
	fmt.Println(class)

	// create multiple variables
	name, age := "Dung", 23
	fmt.Println(name, age)
	fmt.Printf("%s\n", name)

	// type the variable: %T
	fmt.Printf("%T\n", name)

	// const
	const PI = math.Pi
	fmt.Println(PI)

	// Data type: int, float, complex, string, bool, byte, rune
	// int/uint, int8/uint8, int16/uint16, int32/uint32, int64/uint64
	// float32, float64 -> default float64
	
}
