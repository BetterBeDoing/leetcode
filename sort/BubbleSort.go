package sort

func BubbleSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if input[i] > input[j] {
				tmp := input[i]
				input[i] = input[j]
				input[j] = tmp
			}
		}
	}
	return input
}
