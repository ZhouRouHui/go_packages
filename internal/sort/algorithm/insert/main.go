package main

import "fmt"

/**
插入排序
*/

var s []int = []int{2, 5, 3, 9, 4, 6}

func main() {
	l := len(s)
	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[j+1] < s[j] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
	fmt.Println("插入排序的结果：")
	fmt.Println(s)
}
