package main

import "sort"

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if i < j && nums[i]%2 == 0 {
			i++
		}
		if i < j && nums[j]%2 == 1 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func sortArrayByParity3(nums []int) []int {
	even := make([]int, 0, len(nums))
	odd := make([]int, 0, len(nums))
	for _, n := range nums {
		if n%2 == 0 {
			even = append(even, n)
			continue
		}
		odd = append(odd, n)
	}
	return append(even, odd...)
}

func sortArrayByParity2(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i]%2 == 0 && nums[j]%2 != 0 {
			return true
		} else if nums[i]%2 != 0 && nums[j]%2 == 0 {
			return false
		}
		return nums[i] < nums[j]
	})
	return nums
}
