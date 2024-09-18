package leetcode

func spiralOrder(matrix [][]int) []int {
	maxCol := len(matrix[0])
	maxRow := len(matrix)
	sum := maxCol * maxRow

	minRow := -1
	minCol := -1
	res := make([]int, sum)
	row := 0
	col := 0
	moveFlag := 0
	for i := 0; i < sum; i++ {
		res[i] = matrix[row][col]
		row, col = moveNext(row, col, moveFlag%4)
		if col >= maxCol {
			row, col = moveLeft(moveDown(row, col))
			moveFlag++
			minRow++
		}
		if row >= maxRow {
			row, col = moveLeft(moveUp(row, col))
			moveFlag++
			maxCol--
		}
		if col <= minCol {
			row, col = moveRight(moveUp(row, col))
			moveFlag++
			maxRow--
		}
		if row <= minRow {
			row, col = moveRight(moveDown(row, col))
			moveFlag++
			minCol++
		}
	}
	return res
}

func moveNext(x, y, moveFlag int) (_, _ int) {
	switch moveFlag {
	case 0:
		return moveRight(x, y)
	case 1:
		return moveDown(x, y)
	case 2:
		return moveLeft(x, y)
	case 3:
		return moveUp(x, y)
	}
	return x, y
}

func moveRight(x, y int) (_, _ int) {
	y++
	return x, y
}

func moveLeft(x, y int) (_, _ int) {
	y--
	return x, y
}

func moveUp(x, y int) (_, _ int) {
	x--
	return x, y
}

func moveDown(x, y int) (_, _ int) {
	x++
	return x, y
}
