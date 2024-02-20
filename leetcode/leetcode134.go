package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	startPosition := -1
	min := 10000
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum <= min {
			min = sum
			startPosition = i
		}
	}

	if sum >= 0 {
		return (startPosition + 1) % len(gas)
	}
	return -1
}
