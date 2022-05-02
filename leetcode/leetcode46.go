package leetcode

/*Permute
用回溯法解题通常包含以下3个步骤：
1. 针对所给问题，定义问题的解空间
2. 确定易于搜索的解空间结构
3. 以深度优先方式所搜解空间，并在搜索过程中用剪枝函数避免无效搜索。
*/
func Permute(nums []int) (ans [][]int) {
	var res []int
	res = []int{}
	ans = make([][]int, 0)
	var flag []bool
	flag = make([]bool, len(nums))
	//backtrack 传入的l代表的是全排列的深度
	var backtrack func(l int)
	backtrack = func(l int) {
		if l == len(nums) {
			tmp := make([]int, len(nums))
			for i := range res {
				tmp[i] = res[i]
			}
			ans = append(ans, tmp)
			return
		} else {
			for index := range flag {
				if flag[index] == false {
					flag[index] = true
					res = append(res, nums[index])
					backtrack(l + 1)
					flag[index] = false
					res = res[0 : len(res)-1]
				}
			}
		}
	}
	backtrack(0)
	return ans
}
