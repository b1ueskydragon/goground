package p0643

// find the contiguous subarray of given length k that has the maximum average value
func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}
	candidate := maxSum

	for i := k; i < len(nums); i++ {
		candidate = candidate + nums[i] - nums[i-k]

		if maxSum < candidate {
			maxSum = candidate
		}
	}

	return float64(maxSum) / float64(k)
}
