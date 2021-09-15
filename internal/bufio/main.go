package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// https://www.cnblogs.com/zhouqi666/p/9653711.html
/**
bufio 包实现了带缓存的 I/O 操作
它封装一个 io.Reader 或 io.Write 对象，使其具有缓存和一些文本读写功能
*/

// 获取一个测试用的 reader
var r io.Reader = strings.NewReader("abcdefghijklmnopqrstuvwxyz")
var br *bufio.Reader

// 获取一个测试用的 bufio.Writer
var w io.Writer = bytes.NewBuffer(make([]byte, 0))
var bw *bufio.Writer

func init() {
	br = bufio.NewReader(r)
	bw = bufio.NewWriter(w)
}

func main() {

	// bufio 的 Reader 相关
	//aboutReader()

	// bufio 的 Writer 相关
	//aboutWrite()

	// 聚合了 bufio.Reader 和 bufio.Writer 的一个类型
	testReadWriter()
}
