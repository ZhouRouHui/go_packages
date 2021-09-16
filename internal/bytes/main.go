package main

/**
bytes 包 https://www.cnblogs.com/golove/p/3287729.html

对于传入 []byte 的函数，都不会修改传入的参数，返回值要么是参数的副本，要么是参数的切片
*/

func main() {
	// bytes 相关
	aboutBytes()

	// bytes.Reader 相关
	aboutBytesReader()

	// bytes.Buffer 相关
	aboutBytesBuffer()
}
