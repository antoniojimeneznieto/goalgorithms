package main

import (
	"fmt"
	"time"
)

func main() {
	/* Arrays */
	// int32 is 4 bytes, 4 * 5 = 20 bytes total.
	var a [5]int32 // array declaration, default value is 0
	fmt.Println(a)
	fmt.Println(len(a)) // length of the array
	fmt.Println(a[0])   // accessing the first element of the array
	fmt.Println(a[1:])  // slicing the array, last index is exclusive

	fmt.Println(&a[0]) // address of the first element of the array
	fmt.Println(&a[1]) // address of the second element of the array
	fmt.Println(&a[2]) // address of the third element of the array
	fmt.Println(&a[3]) // address of the fourth element of the array
	fmt.Println(&a[4]) // address of the fifth element of the array

	var b [3]int32 = [3]int32{1, 2, 3} // array declaration with initialization
	fmt.Println(b)

	/* Slices */
	var c []int32 = []int32{1, 2, 3} // slice declaration with initialization
	fmt.Printf("Length: %d, Capacity: %d, Value: %v\n", len(c), cap(c), c)

	c = append(c, 4) // append an element to the slice
	fmt.Printf("Length: %d, Capacity: %d, Value: %v\n", len(c), cap(c), c)

	var slice2 []int32 = []int32{5, 6, 7}
	c = append(c, slice2...) // append an slice to the slice
	fmt.Printf("Length: %d, Capacity: %d, Value: %v\n", len(c), cap(c), c)

	var sliceMake []int32 = make([]int32, 3, 8) // make function to create a slice of length 3 and capacity 8
	fmt.Printf("Length: %d, Capacity: %d, Value: %v\n", len(sliceMake), cap(sliceMake), sliceMake)

	/* Maps */
	var m1 map[string]int = make(map[string]int) // map declaration
	fmt.Println(m1)
	var m2 map[string]int = map[string]int{"a": 1, "b": 2} // map declaration with initialization
	fmt.Println(m2)
	fmt.Println(m2["a"]) // accessing the value of the key "a"

	var num, ok = m2["c"] // accessing the value of the key "c" and checking if the key exists
	fmt.Printf("Value: %d, Key exists: %t\n", num, ok)
	num, ok = m2["b"]
	fmt.Printf("Value: %d, Key exists: %t\n", num, ok)

	delete(m2, "b") // delete a key from the map
	fmt.Println(m2)

	/* loops */

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, v := range []int{1, 2, 3} {
		fmt.Println(i, v)
	}

	for k, v := range map[string]int{"a": 1, "b": 2} {
		fmt.Println(k, v)
	}

	/* time */
	var n int = 1000000
	var slice = make([]int, 0, n) // preallocate the slice with the capacity of n
	fmt.Println(timeLoop(slice, n))
	var sliceWithout = []int{} // without preallocating the slice
	fmt.Println(timeLoop(sliceWithout, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var initialTime = time.Now()
	for i := 0; i < n; i++ {
		slice = append(slice, n)
	}
	return time.Since(initialTime)
}
