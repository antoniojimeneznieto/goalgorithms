package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/* variable declaration and printing */
	var x int = 10
	var y, z int = 20, 30 // multiple variable declaration

	println(x, y, z)
	fmt.Println(x, y, z)
	fmt.Printf("Variable a of type: %T has value: %v\n", x, x) // %T is used to print the type of the variable

	a := 10        // shorthand declaration (type inference)
	b, c := 20, 30 // multiple variable declaration with shorthand

	fmt.Println(a, b, c) // fmt.Println is used to print formatted output

	var myString string = "Hello,\nWorld!"
	var MyString2 string = `Hello,World!` // raw string literal

	fmt.Println(myString + MyString2)

	var aBool bool = true
	fmt.Printf("Variable aBool of type: %T has value: %v\n", aBool, aBool)

	/* X bit variables */
	var aFloat float64 = 169123.13123
	var aFloat32 float32 = 169123.13123 // float32 is less precise than float64

	fmt.Printf("Variable aFloat of type: %T has value: %v\n", aFloat, aFloat)
	fmt.Printf("Variable aFloat32 of type: %T has value: %v\n", aFloat32, aFloat32)

	var int32Var int32 = 10
	var int64Var int64 = 10
	fmt.Printf("Variable int32 has value: %v and int64 has value: %v\n", int32Var, int64Var)

	var aComplex complex64 = 1 + 2i
	fmt.Printf("Variable aComplex of type: %T has value: %v\n", aComplex, aComplex)

	/* counting the number of characters */
	var normalString string = "test"
	var specialCharacter string = "æ"

	fmt.Println(len(normalString))
	fmt.Println(len(specialCharacter)) // len returns the number of bytes in a string

	fmt.Println(utf8.RuneCountInString(specialCharacter)) // To get the number of characters use UTF8

	/* float + int */
	var floatVar float64 = 10.6
	var intVar int = 10
	fmt.Println(floatVar + float64(intVar)) // Type conversion
	fmt.Println(int(floatVar) + intVar)     // Type conversion

	// constant values
	const pi float64 = 3.14159
	fmt.Println(pi)
}
