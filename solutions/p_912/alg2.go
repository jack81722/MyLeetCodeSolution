package main

// this solution will be wrong in leetcode
// but it is correct in local environment
func sortArray2(nums []int) []int {
	quick2(nums, 0, len(nums)-1)
	return nums
}

func quick2(nums []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r + 1) / 2
	i, j := l, r
	for i != j {
		if nums[i] <= nums[m] && i < m {
			i++
		}
		if nums[j] >= nums[m] && j > m {
			j--
		}
		if nums[i] > nums[m] || nums[j] < nums[m] {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if r != m {
		quick2(nums, l, m)
	}
	if l != m {
		quick2(nums, m, r)
	}
}
