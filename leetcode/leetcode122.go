package leetcode

func maxProfit(prices []int) int {
	//it's a typical dp problem
	//profit[i] = max(profit[i-1], profit[i-1] + prices[i] - prices[i-1])

	var (
		profit []int = make([]int, len(prices))
	)

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			profit[i] = 0
		} else {
			profit[i] = max(profit[i-1], profit[i-1]+prices[i]-prices[i-1])
		}
	}

	return profit[len(prices)-1]
}
