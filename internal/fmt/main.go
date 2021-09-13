package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// 打印内容，返回值是正确打印内容的数量
	printRes, _ := fmt.Print("a", "b")
	fmt.Printf("printRes = %d\n", printRes)

	// 打印内容
	fmt.Println("hello", "world")

	// 格式化打印
	n := 999
	fmt.Printf("n = %d\n", n)

	// 创建一个 err 信息
	err := fmt.Errorf("errorf func test: %v", errors.New("这是一个错误信息"))
	fmt.Printf("err = %v\n", err)

	/**
	Fprint 系列，返回写入的数量，中文字符一个字算 3 个长度
	*/
	// Fprint 将内容写入资源句柄
	fp, _ := os.OpenFile("./fprint.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer fp.Close()
	//fPrintRes, _ := fmt.Fprint(fp, "a", "b", "c")
	//fmt.Printf("fPrintRes = %d\n", fPrintRes)

	// 同上，带空格，带换行符
	fPrintLnRes, _ := fmt.Fprintln(fp, "d", "e", "f")
	fmt.Printf("fPrintLnRes = %d\n", fPrintLnRes)

	// 格式化的写入资源句柄，
	fPrintfRes, _ := fmt.Fprintf(fp, "这是一个测试数据，%d\n", n)
	fmt.Printf("fPrintfRes = %d\n", fPrintfRes)

	/**
	Sprint 系列，生成字符串
	*/
	sprintRes := fmt.Sprint("a", "b", "c")
	fmt.Printf("sprintRes = %s\n", sprintRes)

	sprintLnRes := fmt.Sprintln("a", "b", "c")
	fmt.Printf("sprintLnRes = %s", sprintLnRes)

	sprintfRes := fmt.Sprintf("this is Sprintf test, %d", n)
	fmt.Printf("sprintfRes = %s\n", sprintfRes)

	/**
	Scan 从标准输入 os.Stdin 读取文本(从终端获取数据)
	从标准输入中扫描文本，读取由空白符分隔的值保存到函数的参数中，标准输入中的换行符视为空白符，返回扫描到的个数
	*/
	var (
		name  string
		age   int
		job   string
		hobby string
	)
	fmt.Printf("please input your name\n") // 终端输出这句话后可以在终端输入内容， Scan 可以获取到对应的内容
	fmt.Scan(&name)

	// Scanln 在遇到换行符时会立刻停止扫描，
	// TODO 测试未通过
	//fmt.Printf("please input your age")
	//fmt.Scanln(&age)

	fmt.Printf("please input your job and hobby\n")
	fmt.Scanf("job = %s, hobby = %s\n", &job, &hobby)
	fmt.Printf("name = %s, age = %d\n", name, age)
}
