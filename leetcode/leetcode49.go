package leetcode

func GroupAnagrams(strs []string) [][]string {
	mem := make(map[string][]string, 0)

	for index := range strs {
		str := strs[index]
		resStr := sort49(str)
		val, ok := mem[resStr]
		if !ok {
			val = make([]string, 0)
			val = append(val, str)
			mem[resStr] = val
		} else {
			val = append(val, str)
			mem[resStr] = val
		}
	}

	res := make([][]string, 0)
	for index := range mem {
		res = append(res, mem[index])
	}
	return res
}

func sort49(str string) string {
	tmp := make([]int, 26)
	aAnc := 'a'
	for index := range str {
		tmp[str[index]-byte(aAnc)]++
	}

	res := ""
	for index := 0; index < 26; index++ {
		if tmp[index] > 0 {
			c := string(uint8(index) + byte(aAnc))
			tmp[index]--
			index--
			res += c
		}
	}
	return res
}
