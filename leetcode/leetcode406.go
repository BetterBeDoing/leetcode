package leetcode

func reconstructQueue(people [][]int) [][]int {
	l := len(people)
	QuickSortByH(people, 0, l-1)
	return nil
}

func QuickSortByH(people [][]int, left int, right int) {
	if left > right {
		return
	}

	j := QuickSortPart(people, left, right)
	QuickSortByH(people, left, j-1)
	QuickSortByH(people, j+1, right)
}

func QuickSortPart(people [][]int, left int, right int) int {
	i, j := left+1, right
	for true {
		for people[i][0] >= people[left][0] {
			if people[i][0] == people[left][0] && people[i][1] >= people[left][1] {
				break
			}
			i++
			if i == right {
				break
			}
		}
		for people[j][0] <= people[left][0] {
			if people[j][0] == people[left][0] && people[j][1] <= people[left][1] {
				break
			}
			j--
			if j == left {
				break
			}
		}
		if i >= j {
			break
		}
		QuickSortExchange(people, i, j)
	}
	QuickSortExchange(people, left, j)
	return j
}

func QuickSortExchange(people [][]int, i int, j int) {
	temp := people[i]
	people[i] = people[j]
	people[j] = temp
}
