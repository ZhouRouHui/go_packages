package main

import (
	"fmt"
	"io/fs"
)

// 文件系统相关接口
func main() {

	// 校验一个路径是否正确
	// todo 不同操作系统的路径分隔符不一致如何解决？
	validPathRes := fs.ValidPath(".\\test.txt")
	fmt.Printf("validPathRes = %t\n", validPathRes)

	// 提供一种方法，支持正则匹配文件名
	//fs.Glob()

	// ReadDir 读取指定的目录并返回一个按文件名排序的目录条目列表。
	//fs.ReadDir()

	// ReadFile从文件系统fs中读取命名的文件并返回其内容。一个成功的调用会返回一个nil错误，而不是io.EOF。
	//fs.ReadFile()
}
