package leetcode

//todo dp[][]可以设计成为一个bool类型的二维数组，然后用动态规划的方式拓展
func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	//回文串分奇数、偶数
	//dp[i][j]其中i代表的是当前长度的回文串，j
	l := len(s)
	dp := make([][]string, l) //最长的情况是s就是回文串

	for i := range dp {
		dp[i] = make([]string, l) //种类最多的情况是l种
	}
	//中心拓展法，遍历dp[i-2][j]种的情况进行拓展
	start := -2
	end := 0

	maxStart := 0
	maxEnd := 0
	for start < l {
		//奇数中心初始化
		if start == -2 {
			for index, _ := range s {
				dp[end][index] = string(s[index])
			}
			start++
			end++
			continue
		}
		//偶数中心初始化
		if start == -1 {
			for index := 0; index < l-1; index++ {
				if string(s[index]) == string(s[index+1]) {
					dp[end][index] = s[index : index+2]
					maxStart = index
					maxEnd = index + 2
				}
			}
			start++
			end++
			continue
		}
		//依据父节点拓展中心
		for j := range dp[start] {
			if dp[start][j] == "" { //当父层没有该中心节点，而且其左右拓展会受阻时直接跳过该层
				continue
			}
			if j < 1 {
				continue
			}
			if j+start+1 >= l {
				continue
			}
			if s[j-1] == s[j+start+1] {
				tmp := s[j-1 : j+start+2]
				dp[end][j-1] = tmp
				maxStart = j - 1
				maxEnd = j + start + 2
			}
		}
		end++
		start++
		if start == end {
			break
		}
	}

	return s[maxStart:maxEnd]
}
