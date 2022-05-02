package leetcode

func numTrees(n int) int {
	res := make([]int, n+1)
	res[0], res[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			res[i] += res[j-1] * res[i-j]
		}
	}
	return res[n]
}
