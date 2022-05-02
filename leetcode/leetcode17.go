package leetcode

var mem = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	combinations = []string{}
	getCombination(digits, 0, "")
	return combinations
}
func getCombination(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		thisNum := string(digits[index])
		letters := mem[thisNum]
		for i := 0; i < len(letters); i++ {
			getCombination(digits, index+1, combination+string(letters[i]))
		}
	}
}
