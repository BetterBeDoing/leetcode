package leetcode

import "math"

func canPartition(nums []int) bool {
	l := len(nums)
	sum, maxNum := getSumAndMax(nums)
	target := sum / 2
	flag := checkInit(sum, maxNum)
	if !flag {
		return flag
	}

	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	for i := 0; i < l; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true

	for i := 1; i < l; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[l-1][target]

}

func getSumAndMax(nums []int) (sum int, maxNum int) {
	sum = 0
	maxNum = math.MinInt
	for _, v := range nums {
		sum += v
		if v > maxNum {
			maxNum = v
		}
	}
	return sum, maxNum
}

func checkInit(sum int, maxNum int) bool {
	if sum%2 == 1 {
		return false
	}
	if maxNum > sum/2 {
		return false
	}
	return true
}
