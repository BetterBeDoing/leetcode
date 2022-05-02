package leetcode

//Subsets
//使用回溯法
//1,2,3和1，3，2认为是相同的组合
func Subsets(nums []int) [][]int {
	//res
	ans := [][]int{[]int{}}
	res := make([]int, 0)

	used := make([]bool, len(nums))
	var backtrack func(l, pos int)
	backtrack = func(l, pos int) {

		for index := range nums {
			if pos > index {
				continue
			}
			if used[index] == false {
				res = append(res, nums[index])
				//存储中间状态
				tmp := make([]int, len(res))
				for index := range res {
					tmp[index] = res[index]
				}
				ans = append(ans, tmp)
				used[index] = true
				backtrack(l+1, index)
				used[index] = false
				res = res[0 : len(res)-1]
			}
		}
	}

	backtrack(0, 0)
	return ans
}
