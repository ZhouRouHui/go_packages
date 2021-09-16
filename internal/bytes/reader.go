package main

import (
	"bytes"
	"fmt"
)

/**
bytes.Reader 结构体方法
*/

func aboutBytesReader() {
	// 实例化一个 bytes.Reader 对象
	// bytes.Reader 实现了如下接口：
	// io.ReadSeeker
	// io.ReaderAt
	// io.WriterTo
	// io.ByteScanner
	// io.RuneScanner
	reader := bytes.NewReader([]byte("hello world"))
	fmt.Printf("reader = %v\n", reader)

	// 返回未读取部分的数据长度
	n := reader.Len()
	fmt.Printf("n = %d\n", n)

	// 返回底层数据的总长度，方便 ReadAt 使用，返回值永远不变。
	sizeRes := reader.Size()
	fmt.Printf("sizeRes = %d\n", sizeRes)

	// 将底层数据切换为 b，同时复位所有标记（读取位置等信息）。
	reader.Reset([]byte("a new byte"))
}
