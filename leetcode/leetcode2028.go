package leetcode

func MissingRolls(rolls []int, mean int, n int) []int {
	missingSum := mean * (n + len(rolls))
	for _, roll := range rolls {
		missingSum -= roll
	}
	if missingSum < n || missingSum > n*6 {
		return nil
	}

	quotient, remainder := missingSum/n, missingSum%n
	ans := make([]int, n)
	for i := range ans {
		ans[i] = quotient
		if i < remainder {
			ans[i]++
		}
	}
	return ans
}
