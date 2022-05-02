package leetcode

import "strconv"

func calPoints(ops []string) int {
	res := []int{}
	sum := 0
	for _, val := range ops {
		switch val {
		case "C":
			res = res[:len(res)-2]
		case "D":
			res = append(res, res[len(res)-1]*2)
		case "+":
			res = append(res, res[len(res)-1]+res[len(res)-2])
		default:
			tmp, _ := strconv.Atoi(val)
			res = append(res, tmp)
		}
	}
	for _, val := range res {
		sum += val
	}
	return sum
}
