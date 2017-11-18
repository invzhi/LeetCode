package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var indices []int

	for i, num := range nums {
		if v, isExist := m[target-num]; isExist {
			indices = []int{i, v}
		} else {
			m[num] = i
		}
	}
	return indices
}
