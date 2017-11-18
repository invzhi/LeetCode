package leetcode

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)

	for i, number := range numbers {
		m[number] = i
	}

	var indices []int

	for i, number := range numbers {
		if j, isExist := m[target-number]; isExist && i != j {
			indices = []int{i + 1, j + 1}
			break
		}
	}
	return indices
}
