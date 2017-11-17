package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, isExist := m[target-num]; isExist {
			return []int{i, v}
		} else {
			m[num] = i
		}
	}
	return nil
}
