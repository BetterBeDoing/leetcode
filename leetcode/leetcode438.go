package leetcode

func FindAnagrams(s string, p string) []int {
	ls := len(s)
	lp := len(p)
	if ls < lp {
		return nil
	}

	var sMem [26]int
	var pMem [26]int
	for i, v := range p {
		sMem[s[i]-'a']++
		pMem[v-'a']++
	}

	res := make([]int, 0)
	if sMem == pMem {
		res = append(res, 0)
	}

	for i, _ := range s[:ls-lp] {
		sMem[s[i]-'a']--
		sMem[s[i+lp]-'a']++

		if sMem == pMem {
			res = append(res, i)
		}
	}

	return res
}
