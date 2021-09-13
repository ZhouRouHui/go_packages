package main

import (
	"fmt"
)

/**
冒泡排序
*/

var s []int = []int{2, 5, 3, 9, 4, 6}

// 把小的往前放
func bubble1(l int) {
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// 把大的往后冒
func bubble2(l int) {
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}

func main() {
	l := len(s)
	//bubble1(l)

	bubble2(l)

	fmt.Printf("冒泡排序结果：\n")
	fmt.Println(s)
}
