package main

import (
	"fmt"
)

/**
选择排序
*/

var s []int = []int{2, 5, 3, 9, 4, 6}

// 常规思路的
func select1(l int) {
	for i := 0; i < l; i++ {
		t := i
		for j := i + 1; j < l; j++ {
			if s[t] > s[j] {
				t = j
			}
		}
		if t != i {
			s[i], s[t] = s[t], s[i]
		}
	}
}

// 优化后的
func select2(l int) {
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func main() {
	l := len(s)
	//select1(l)

	select2(l)

	fmt.Println("选择排序的结果：")
	fmt.Println(s)
}
