// 475. Heaters

// Winter is coming! Your first job during the contest is to design a standard heater with fixed warm radius to warm all the houses.
// Now, you are given positions of houses and heaters on a horizontal line, find out minimum radius of heaters so that all houses could be covered by those heaters.
// So, your input will be the positions of houses and heaters separately, and your expected output will be the minimum radius standard of heaters.

// Note:
// 1. Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
// 2. Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
// 3. As long as a house is in the heaters' warm radius range, it can be warmed.
// 3. All the heaters follow your radius standard and the warm radius will the same.

// Example 1:
// Input: [1,2,3],[2]
// Output: 1
// Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.

// Example 2:
// Input: [1,2,3,4],[1,4]
// Output: 1
// Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.

package leetcode

import "sort"

func findRadius(houses, heaters []int) int {
	sort.Ints(heaters)

	var radius int
	for _, house := range houses {
		i := findIndex(heaters, house)

		min := int(1e9)
		if i >= 0 {
			min = house - heaters[i]
		}
		if i+1 < len(heaters) && heaters[i+1]-house < min {
			min = heaters[i+1] - house
		}

		if min > radius {
			radius = min
		}
	}
	return radius
}

func findIndex(nums []int, num int) int {
	lo, hi := -1, len(nums)-1
	for lo < hi {
		mid := hi - (hi-lo)/2
		if nums[mid] < num {
			lo = mid
		} else if nums[mid] > num {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return lo
}
