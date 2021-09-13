package main

import (
	"fmt"
	"sort"
)

var intSlice []int = []int{2, 4, 1, 3, 22, 3}
var float64Slice []float64 = []float64{2.1, 3.2, 3.1, 5.5}
var stringSlice []string = []string{"zrh", "loedan", "kevin"}

func main() {
	/**
	Int 相关操作
	*/
	// Ints 对 int 的切片进行升序排序
	sort.Ints(intSlice)
	fmt.Printf("sorted intSlice = %v\n", intSlice)

	// IntsAreSorted 判断一个 int 的切片是否是升序排序的
	intsAreSortedRes := sort.IntsAreSorted(intSlice)
	fmt.Printf("intsAreSortedRes = %t\n", intsAreSortedRes)

	// SearchInts 在一个升序排序的整数切片中搜索某个元素，并返回下标值
	// 如果内容不存在，返回 len(slice)
	searchIntsRes := sort.SearchInts(intSlice, 300)
	fmt.Printf("searchIntsRes = %d\n", searchIntsRes)

	/**
	float64 相关操作
	*/
	// Float64s 对 float 64 的切片升序排序
	sort.Float64s(float64Slice)
	fmt.Printf("sorted float64Slice = %v\n", float64Slice)

	// Float64sAreSorted 判断一个 float64 的切片是否是升序排序的
	float64sAreSortedRes := sort.Float64sAreSorted(float64Slice)
	fmt.Printf("float64sAreSortedRes = %t\n", float64sAreSortedRes)

	// SearchFloat64s 在一个升序排序的 float64 切片中搜索某个元素，并返回下标值
	// 如果内容不存在，返回 len(slice)
	searchFloat64sRes := sort.SearchFloat64s(float64Slice, 5.5)
	fmt.Printf("searchFloat64sRes = %d\n", searchFloat64sRes)

	/**
	string 相关操作
	*/
	// Strings 对 string 的切片进行升序排序
	sort.Strings(stringSlice)
	fmt.Printf("sorted stringSlice = %v\n", stringSlice)

	stringsAreSortedRes := sort.StringsAreSorted(stringSlice)
	fmt.Printf("stringsAreSortedRes = %t\n", stringsAreSortedRes)

	// SearchStrings 在一个升序排序的 string 切片中搜索某个元素，并返回下标值
	// 如果内容不存在，返回 len(slice)
	searchStringsRes := sort.SearchStrings(stringSlice, "zrh")
	fmt.Printf("searchStringsRes = %d\n", searchStringsRes)
}
