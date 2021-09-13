package main

import "fmt"

/**
快速排序
*/

var s []int = []int{6, 22, 34, 213, 12, 0, 8, 55}

func quick(s []int) []int {
	l := len(s)
	if l == 1 {
		return s
	}
	poiv := s[0]
	left := []int{}
	right := []int{}

	for i := 1; i < l; i++ {
		if s[i] > poiv {
			right = append(right, s[i])
		}
		if s[i] < poiv {
			left = append(left, s[i])
		}
	}

	if len(left) > 0 {
		left = quick(left)
	}
	if len(right) > 0 {
		right = quick(right)
	}
	res := []int{}
	res = append(res, left...)
	res = append(res, poiv)
	res = append(res, right...)
	return res
}

func main() {
	s = quick(s)
	fmt.Println("快速排序的结果：")
	fmt.Println(s)
}
