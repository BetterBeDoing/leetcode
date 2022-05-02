package leetcode

/*BoolLongestPalindrome
bool[i][j] 代表i:j的子串是否是回文串;
bool[i][i]=true;
bool[i-1][j-1]=if s[i]==s[j]&&bool[i][j];
maxI&&maxJ mark the position;
*/
func BoolLongestPalindrome(s string) string {
	//当s为空或者为1的时候直接返回即可
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}

	maxS := 0
	maxE := 0

	l := len(s)
	//将[][]bool初始化
	mem := make([][]bool, l)
	for i := range mem {
		mem[i] = make([]bool, l)
		mem[i][i] = true
		if i+1 >= l {
			continue
		}
		mem[i][i+1] = s[i] == s[i+1]
		if mem[i][i+1] == true {
			maxS = i
			maxE = i + 1
		}
	}
	//遍历二维数组

	for lg := 2; lg < l; lg++ {
		for index := 0; index < l; index++ {
			if index < 0 {
				continue
			}
			if index+lg >= l {
				break
			}
			f1 := mem[index+1][index+lg-1]
			f2 := s[index] == s[index+lg]
			mem[index][index+lg] = f1 && f2
			if mem[index][index+lg] == true {
				maxS = index
				maxE = index + lg
			}
		}
	}
	return s[maxS : maxE+1]
}
