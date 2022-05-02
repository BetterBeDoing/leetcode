package leetcode

var ans []string

//回溯法
func generateParenthesis(n int) []string {
	ans = []string{}
	s := make([]byte, 0)
	backTrack(s, 0, 0, n)
	return ans
}

func backTrack(s []byte, left int, right int, n int) {
	if len(s) == 2*n {
		ans = append(ans, string(s))
	}
	if left < n {
		s = append(s, '(')
		backTrack(s, left+1, right, n)
		s = s[0 : len(s)-1]
	}
	if right < left {
		s = append(s, ')')
		backTrack(s, left, right+1, n)
		s = s[0 : len(s)-1]
	}
}
