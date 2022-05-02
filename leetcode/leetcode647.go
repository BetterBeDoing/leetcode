package leetcode

func CountSubstrings(s string) int {
	//中心拓展法
	l := len(s)
	sum := 0
	for i := 0; i < l; i++ {
		for j := 0; j < 2; j++ {
			left := i
			right := i + j
			for left >= 0 && right < l && s[left] == s[right] {
				left--
				right++
				sum++
			}
		}
	}
	return sum
}

/*func CountSubstrings(s string) int {
	l := len(s)
	sum := 0
	for t := 0; t < l; t++ {
		for i := 0; i < l; i++ {
			if i+t < l {
				temp := s[i : i+t+1]
				sum += JudgeReverse(temp)
			}
		}
	}
	return sum
}*/

/*func JudgeReverse(s string) int {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-i-1] {
			return 0
		}
	}
	return 1
}*/
