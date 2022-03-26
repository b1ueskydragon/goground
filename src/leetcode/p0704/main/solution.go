package main

// nums is sorted in ascending order
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		midVal := nums[m]
		if midVal == target {
			return m
		}
		if midVal > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
