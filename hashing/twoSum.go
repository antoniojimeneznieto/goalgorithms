package hashing

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
