package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 45

	var t0 = time.Now()
	println(fibMemo(n))
	fmt.Printf("The time taken is %v\n", time.Since(t0))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibMemo(n int) int {
	memo := make([]int, n+1)
	return fibMemoHelper(n, &memo)
}

func fibMemoHelper(n int, memo *[]int) int {
	if n <= 1 {
		return n
	}
	if (*memo)[n] == 0 {
		(*memo)[n] = fibMemoHelper(n-1, memo) + fibMemoHelper(n-2, memo)
	}
	return (*memo)[n]
}
