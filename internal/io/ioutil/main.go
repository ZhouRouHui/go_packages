package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// StringReader 封装一个 Reader 对象
type StringReader struct {
	s      []byte // content
	cursor int    // latest read position
	len    int    // content length
}

// NewStringReader 实例一个 StringReader
func NewStringReader(content string) *StringReader {
	contentByte := []byte(content)
	return &StringReader{s: contentByte, cursor: -1, len: len(contentByte)}
}

// Read 读取对象里面的内容
func (s *StringReader) Read(p []byte) (n int, err error) {
	nextIndex := s.cursor + 1
	lr, lp := len(s.s[nextIndex:]), len(p)

	// 游标已到内容尾部
	if s.cursor == (s.len - 1) {
		return 0, io.EOF
	}

	if lr <= lp { // 剩余可读内容小于暂存区长度，则全量读取
		n = copy(p, s.s[nextIndex:])
		s.cursor = s.len - 1
		return n, nil
	} else { // 剩余可读内容大雨暂存区长度，则部分读取
		n = copy(p, s.s[nextIndex:(nextIndex+lp+1)])
		s.cursor += lp
		return n, nil
	}
}

// Reset 重置 cursor
func (s *StringReader) Reset() {
	s.cursor = -1
}

func main() {
	// ioutil.ReadAll 从一个实现了 io.Reader 接口的对象中读取内容
	var myStringReader *StringReader
	myStringReader = NewStringReader("12345678901")
	readAllRes, err := ioutil.ReadAll(myStringReader)
	fmt.Printf("resAllRes = %s, err = %v\n", string(readAllRes), err)
	fmt.Println()

	readFileRes, _ := ioutil.ReadFile("test.txt")
	//fmt.Println(readFileRes)
	fmt.Printf("readFileRes = %s\n", string(readFileRes))

	// 将内容写进文件
	ioutil.WriteFile("./test.txt", []byte("\nhello world"), os.ModeAppend)

	// 读取一个目录
	readDirRes, _ := ioutil.ReadDir("./")
	fmt.Printf("readDirRes = %v\n", readDirRes)
	for _, v := range readDirRes {
		fmt.Printf("readDirRes.v.Name = %s\n", v.Name())
	}
}
