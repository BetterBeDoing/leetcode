package leetcode

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	l := 0
	flag := false
	for i := last; i > -1; i-- {
		if s[i] != ' ' {
			flag = true
			l++
		} else if s[i] == ' ' && flag == true {
			break
		}
	}
	return l
}
