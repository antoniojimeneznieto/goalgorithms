package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string = "Résumé"
	var indexed = myString[1]
	/*
		A character in Go is represented by a byte (following utf8). A byte is an alias for uint8,
		If it is a special character (like é), it will be represented by more than one byte.
		When we try to access the character at index 1, we get the byte at that index, however it is not the character.
		If we iterate over the string, we will get the characters.
	*/

	fmt.Printf("Indexed type: %T and indexed value: %v\n", indexed, indexed)
	for i, c := range myString {
		fmt.Printf("Character at index %d is %v and the corresponding character is %c\n", i, c, c)
	}

	// For this reason, when we do len(myString), we get the number of bytes in the string, not the number of characters.
	fmt.Println(len(myString))

	/* Rune */
	// A rune is an alias for int32 and it represents a Unicode code point.
	var myStringAsRune = []rune("Résumé")
	var indexedRune = myStringAsRune[1]

	fmt.Printf("Indexed type: %T and indexed value: %v\n", indexedRune, indexedRune)
	fmt.Println(len(myStringAsRune))
	for i, c := range myStringAsRune {
		fmt.Printf("Character at index %d is %v and the corresponding character is %c\n", i, c, c)
	}

	// string as slice
	var myStringAsSlice = []string{"r", "e", "s", "u", "m", "é"}
	fmt.Println(myStringAsSlice)
	var strBuilder strings.Builder
	for _, c := range myStringAsSlice {
		strBuilder.WriteString(c)
	}
	var str string = strBuilder.String()
	fmt.Println(str)

}
