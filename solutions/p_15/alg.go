package main

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0, 100)
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		two := twoSum(nums, -nums[i], i+1)
		for _, arr := range two {
			seg := []int{
				nums[i],
				nums[arr[0]],
				nums[arr[1]],
			}
			sort.Slice(seg, func(x, y int) bool {
				return seg[x] < seg[y]
			})
			result = append(result, seg)
		}
	}
	return result
}

func twoSum(nums []int, target int, offset int) [][]int {
	result := make([][]int, 0, 100)
	set := Set{}
	remain := map[int]int{}
	for i := offset; i < len(nums); i++ {
		if set[nums[i]] {
			continue
		}
		if _, ok := remain[nums[i]]; ok {
			set[nums[i]] = true
			result = append(result, []int{remain[nums[i]], i})
		}
		sub := target - nums[i]
		remain[sub] = i
	}
	return result
}

type Set map[int]bool
