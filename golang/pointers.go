package main

import "fmt"

func main() {
	var a int8 = 10
	var b *int8 = new(int8) // Pointer to int8

	*b = 20 // Dereferencing the pointer b and assigning 20 to the value it points to
	fmt.Printf("The addresses of a and b are %p and %p and the values are %v and %v\n", &a, b, a, *b)
	b = &a // Assigning the address of a to b
	fmt.Printf("The addresses of a and b are %p and %p and the values are %v and %v\n", &a, b, a, *b)

	// Slices contains pointers
	var c []int = []int{1, 2, 3}
	var d []int = c // d is a slice that points to the same underlying array as c

	d[0] = 10
	fmt.Println("the value of c[0] is", c[0], "and the value of d[0] is", d[0])
}
