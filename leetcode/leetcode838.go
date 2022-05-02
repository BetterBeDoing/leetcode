package leetcode

import "bytes"

func pushDominoes(dominoes string) string {
	s := []byte(dominoes)
	i, n, left := 0, len(s), byte('L')
	for i < n {
		j := i
		for j < n && s[j] == '.' {
			j++
		}
		right := byte('R')
		if j < n {
			right = s[j]
		}
		if left == right {
			for i < j {
				s[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' {
			k := j - 1
			for i < k {
				s[i] = 'R'
				s[k] = 'L'
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	return string(s)
}

func dpPushDominoes(dominoes string) string {
	n := len(dominoes)
	q := []int{}
	time := make([]int, n)
	for i := range time {
		time[i] = -1
	}
	force := make([][]byte, n)
	for i, ch := range dominoes {
		if ch != '.' {
			q = append(q, i)
			time[i] = 0
			force[i] = append(force[i], byte(ch))
		}
	}

	ans := bytes.Repeat([]byte{'.'}, n)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if len(force[i]) > 1 {
			continue
		}
		f := force[i][0]
		ans[i] = f
		ni := i - 1
		if f == 'R' {
			ni = i + 1
		}
		if 0 <= ni && ni < n {
			t := time[i]
			if time[ni] == -1 {
				q = append(q, ni)
				time[ni] = t + 1
				force[ni] = append(force[ni], f)
			} else if time[ni] == t+1 {
				force[ni] = append(force[ni], f)
			}
		}
	}
	return string(ans)
}
