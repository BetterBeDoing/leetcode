package leetcode

func isSubsequence(s string, t string) bool {
	// Check if s is a subsequence of t
	// ids-index of s;idt-index of t
	ids := 0
	idt := 0
	for ids < len(s) && idt < len(t) {
		if s[ids] == t[idt] {
			ids++
		}
		idt++
	}

	return ids == len(s)
}
