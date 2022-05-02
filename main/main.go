package main

import (
	"fmt"
)

func groupAnagrams(strs []string) {
	//tmpStrs:=strs
	tmpStrs := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		tmpStrs[i] = strs[i]
	}

	for i := 0; i < len(tmpStrs); i++ {
		//当前的字符串
		tmpStr := tmpStrs[i]
		res := sort(tmpStr)
		//fmt.Println(res)
		tmpStrs[i] = res
	}

	//fmt.Println(tmpStrs)
	resMap := make(map[string][]string, len(strs))
	for i := 0; i < len(strs); i++ {
		val, ok := resMap[tmpStrs[i]]
		if ok {
			val = append(val, strs[i])
			resMap[tmpStrs[i]] = val
		} else {
			tmp := make([]string, 1)
			tmp[0] = strs[i]
			resMap[tmpStrs[i]] = tmp
		}
	}

	fmt.Println(resMap)
}

func sort(str string) string {
	anchor := byte('a')
	tmpMem := make([]int, 26)

	for j := 0; j < len(str); j++ {
		origin := byte(str[j])
		tmpMem[origin-anchor]++
	}

	var newStr = make([]byte, len(str))
	for j := 0; j < len(tmpMem); {
		if tmpMem[j] > 0 {
			newStr = append(newStr, anchor+byte(j))
			tmpMem[j]--
		} else {
			j++
		}
	}
	return string(newStr)
}

/*var x = make(chan int, 1)

func consumer() {
	select {
	case a := <-x:
		fmt.Println("c:", a)
	default:
		fmt.Println("c is waiting")
	}
}
func producer() {
	a := rand.Int()
	select {
	case x <- a:
		fmt.Println("p:", a)
	default:
		fmt.Println("p is waiting")
	}
}*/
//两个栈实现队列
//in，out，count
type stack struct {
	elem          []int
	end, len, cap int
}

func (s *stack) init(cap int) {
	s.elem = make([]int, cap)
	s.end = 0
	s.len = 0
	s.cap = cap
}

func (s *stack) push(input int) {
	s.elem[s.end] = input
	s.len++
	s.end++
}

func (s *stack) pop() int {
	s.len--
	s.end--
	res := s.elem[s.end]
	s.elem[s.end] = -1
	return res
}

func (s stack) getCap() int {
	return s.cap
}

func (s stack) getLen() int {
	return s.len
}

type queue struct {
	stack0, tmp stack
	cap         int
	len         int
}

func (q *queue) init(cap int) {
	q.cap = cap
	q.stack0.init(q.cap)
	q.tmp.init(q.cap)
	q.len = q.stack0.getLen()
}

func (q *queue) inQueue(input int) {
	q.len++
	q.stack0.push(input)
}

func (q *queue) outQueue() int {
	for q.stack0.len > 0 {
		elem := q.stack0.pop()
		q.tmp.push(elem)
	}
	res := q.tmp.pop()
	for q.tmp.len > 0 {
		elem := q.tmp.pop()
		q.stack0.push(elem)
	}
	q.len--
	return res
}

func (q queue) getLen() int {
	return q.stack0.len
}
func main() {
	x := queue{}
	x.init(11)
	for i := 0; i < 10; i++ {
		x.inQueue(i)
	}
	fmt.Println(x)

	for {
		fmt.Println(x.outQueue())
	}

	fmt.Println(x)
}
