package leetcode

func reverseWords(s string) string {
	ns := []byte(s)
	ns = removeRepeatedSpacebar(ns)
	if len(ns) < 2 {
		return string(ns)
	}
	reverseChar(ns)
	reverseWord(ns)
	return string(ns)
}

/*
reverse each char in the string
*/
func reverseChar(input []byte) {
	for i := 0; i < len(input)/2; i++ {
		exchageChar(&input[i], &input[len(input)-1-i])
	}
}

func exchageChar(input0, input1 *byte) {
	tmp := *input0
	*input0 = *input1
	*input1 = tmp
}

/*
reverse each word in the string
example:
the sky is blue
eht yks si eulb
*/
func reverseWord(input []byte) {
	slow := 0
	fast := slow
	for fast < len(input) {
		for ; fast < len(input) && input[fast] != ' '; fast++ {
		}
		for i := slow; i < (fast+slow)/2; i++ {
			exchageChar(&input[i], &input[fast+slow-i-1])
		}
		slow = fast + 1
		fast = slow
	}
}

func removeRepeatedSpacebar(input []byte) []byte {
	start := 0
	for ; input[start] == ' '; start++ {
	}
	end := len(input) - 1
	for ; input[end] == ' '; end-- {

	}
	end++
	for i := start; i < end; {
		if input[i] == ' ' {
			j := i
			for ; input[j] == ' '; j++ {
			}
			num := j - i
			if num == 1 {
				i++
				continue
			} else {
				copy(input[i+1:end-num+1], input[j:end])
				i++
				end -= num - 1
			}

		} else {
			i++
		}
	}
	return input[start:end]
}
