package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 22

	fmt.Println(twoSum(&nums, target))
	fmt.Println(twoSumIterating(&nums, target))
}

func twoSum(nums *[]int, target int) []int {
	m := make(map[int]int)
	for i, v := range *nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

func twoSumIterating(nums *[]int, target int) []int {
	for i, v := range *nums {
		for j, w := range (*nums)[i+1:] {
			if v+w == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return nil
}
