package main

import (
	"fmt"
	"io"
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

// ReadFuncTest Read 相关函数测试
func ReadFuncTest() {
	// 5 个 byte 容量的暂存区，用于存放读取到的内容
	strTmp := make([]byte, 5)
	var err error
	var n int
	var n64 int64

	// 遵循 io.Reader 接口
	myStrReader := NewStringReader("my string reader")

	n, err = myStrReader.Read(strTmp)
	fmt.Printf("myStrReader.Read func test: strTmp[:n] = %s, n = %d, err = %v\n", strTmp[:n], n, err) // my st 5 <nil>

	n, err = myStrReader.Read(strTmp)
	fmt.Printf("myStrReader.Read func test: strTmp[:n] = %s, n = %d, err = %v\n", strTmp[:n], n, err) // ring  5 <nil>

	n, err = myStrReader.Read(strTmp)
	fmt.Printf("myStrReader.Read func test: strTmp[:n] = %s, n = %d, err = %v\n", strTmp[:n], n, err) // reade 5 <nil>

	n, err = myStrReader.Read(strTmp)
	fmt.Printf("myStrReader.Read func test: strTmp[:n] = %s, n = %d, err = %v\n", strTmp[:n], n, err) // r 1 <nil>

	myStrReader.Reset()

	// 读取所有内容
	readAllRes, err := io.ReadAll(myStrReader)
	fmt.Printf("readAllRes = %s, err = %v\n", string(readAllRes), err)

	// 断言最少读
	// io.ReadAtLeast 贪婪读，至少读 min 个即视为成功，尽可能的读 len(buf)
	// 当读取的内容字节数 n == 0, err = io.EOF
	// 当 0 < n < min, err = io.ErrUnexpectedEOF
	// 当 n >= min, err = nil
	n, err = io.ReadAtLeast(myStrReader, strTmp, 3)
	fmt.Printf("io.ReadAtLeast func test: strTmp[:n] = %s, n = %d, err = %v\n", strTmp[:n], n, err)

	// 断言全量读
	// io.ReadFull 断言读，必须读 len(buf) 才视为成功
	// 当读取的内容字节数 n == 0, err = io.EOF
	// 当 0 < n < len(buf), err = io.ErrUnexpectedEOF
	// 当 n == len(buf), err = nil
	n, err = io.ReadFull(myStrReader, strTmp)
	fmt.Printf("io.ReadFull func test: strTmp[:n] = %s, n = %d, err = %v\n", strTmp[:n], n, err)

	/**
	Copy 系列函数
	*/
	// 从参数二中将内容复制到参数一中，直到EOF或者发生错误，返回复制的字节数和遇到的第一个 err
	n64, err = io.Copy(os.Stdout, myStrReader)
	fmt.Println()
	fmt.Printf("io.Copy func test: n64 = %d, err = %v\n", n64, err)

	// 从参数二中 copy 参数三个长度的内容到参数一
	n64, err = io.CopyN(os.Stdout, myStrReader, 5)
	fmt.Println()
	fmt.Printf("io.CopyN func test: n64 = %d, err = %v\n", n64, err)

}
