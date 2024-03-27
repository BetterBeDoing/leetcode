package leetcode

func LongestCommonPrefix(strs []string) string {
	var res string

	if len(strs) == 0 {
		return ""
	}

	res = strs[0]

	if len(strs) == 1 {
		return res
	}

	for i := 1; i < len(strs); i++ {
		res = compareString(res, strs[i])
	}
	return res
}

func compareString(res, strs string) string {
	l := min(len(res), len(strs))
	stop := 0
	for ; stop < l; stop++ {
		if res[stop] != strs[stop] {
			break
		}
	}
	return res[0:stop]
}
