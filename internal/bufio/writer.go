package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// bufio.Writer 实现了为 io.Writer 接口对象提供缓冲。
// 如果在向一个 Writer 类型值写入时遇到错误，该对象将不再接受任何数据，返回该错误数据都写入后，
// 调用者有义务调用 Flush 方法，保证所有的数据都交给了下层的 io.Writer
//
// 创建一个默认的 bufio.Writer，默认的大小是 4096
func testNewWriter() {
	b := bytes.NewBuffer(make([]byte, 0))
	newWriterObj := bufio.NewWriter(b)

	// Available() 返回缓冲中还有多少字节未使用
	// Buffered() 返回缓冲中已使用的字节数
	fmt.Printf("availableRes = %d, bufferedRes = %d\n", newWriterObj.Available(), newWriterObj.Buffered()) // availableRes = 4096, bufferedRes = 0

	// 写入数据
	newWriterObj.WriteString("zrh")

	// Available() 返回缓冲中还有多少字节未使用
	// Buffered() 返回缓冲中已使用的字节数
	fmt.Printf("availableRes = %d, bufferedRes = %d\n", newWriterObj.Available(), newWriterObj.Buffered()) // availableRes = 4093, bufferedRes = 3

	// 丢弃缓冲中的数据，清楚人和错误，将 b 重置为将其输出写入 w
	newWriterObj.Reset(b)

	// Available() 返回缓冲中还有多少字节未使用
	// Buffered() 返回缓冲中已使用的字节数
	fmt.Printf("availableRes = %d, bufferedRes = %d\n", newWriterObj.Available(), newWriterObj.Buffered()) // availableRes = 4096, bufferedRes = 0
}

// 创建一个自定义长度的 bufio.Writer
func testNewWriterSize() {
	b := bytes.NewBuffer(make([]byte, 0))
	newWriterSizeObj := bufio.NewWriterSize(b, 10)

	// Available() 返回缓冲中还有多少字节未使用
	// Buffered() 返回缓冲中已使用的字节数
	fmt.Printf("availableRes = %d, bufferedRes = %d\n", newWriterSizeObj.Available(), newWriterSizeObj.Buffered()) // availableRes = 10, bufferedRes = 0

	// 写入数据
	newWriterSizeObj.WriteString("zrh")

	// Available() 返回缓冲中还有多少字节未使用
	// Buffered() 返回缓冲中已使用的字节数
	fmt.Printf("availableRes = %d, bufferedRes = %d\n", newWriterSizeObj.Available(), newWriterSizeObj.Buffered()) // availableRes = 7, bufferedRes = 3

	// 丢弃缓冲中的数据，清楚人和错误，将 b 重置为将其输出写入 w
	newWriterSizeObj.Reset(b)

	// Available() 返回缓冲中还有多少字节未使用
	// Buffered() 返回缓冲中已使用的字节数
	fmt.Printf("availableRes = %d, bufferedRes = %d\n", newWriterSizeObj.Available(), newWriterSizeObj.Buffered()) // availableRes = 10, bufferedRes = 0
}

// func (b *Writer) Write(p []byte) (nn int, err error)
// Write 将 p 中的数据写入 b 中，返回写入的字节数。
// 如果写入的字节数小于 p 的长度，则返回一个错误信息
//
// func (b *Writer) Flush() error
// Flush 将缓存中的数据提交到底层的 io.Writer 中，也就是这里示例中的 w
func testWrite() {
	p := []byte{'a', 'b', 'c'}
	n, err := bw.Write(p)
	fmt.Printf("写入 bw 的长度 = %d, err = %v\n", n, err)

	fmt.Printf("%q\n", w) // ""
	// 将缓存的数据写入 w 中
	bw.Flush()
	fmt.Printf("%q\n", w)

	// 不影响 p 原来的内容
	fmt.Printf("p = %c\n", p) // "abc"
}

// 功能同 Write，只不过以字符串形式写入
func testWriteString() {
	s := "kevin"
	n, err := bw.WriteString(s)
	fmt.Printf("n = %d, err = %v\n", n, err)

	fmt.Printf("%q\n", w) // ""
	// 将缓存的数据写入 w 中
	bw.Flush()
	fmt.Printf("%q\n", w) // "kevin"
}

// 写入单个字节
func testWriteByte() {
	err := bw.WriteByte('z')
	fmt.Printf("err = %v\n", err)

	bw.Flush()
	fmt.Printf("w = %q\n", w)
}

// 写入单个 rune，支持中文
func testWriteRune() {
	size, err := bw.WriteRune('文')
	fmt.Printf("size = %d, err = %v\n", size, err)

	bw.Flush()
	fmt.Printf("w = %q\n", w)
}

// ReadFrom 实现了 io.ReaderFrom 接口
// 从一个 io.Reader 的缓冲中读取数据并写入到当前 bw 对象中，
// 这里不需要手动执行 Flush，内部会将数据写入到底层的 io.Writer 中，也就是 w 中
//
// 注意，此时这个 io.Reader, 也就是本例中的 br 全部内容已被读出，再读就返回 EOF 错误
func testReadFrom() {
	n, err := bw.ReadFrom(br)
	fmt.Printf("n = %d, err = %v\n", n, err) // n = 26, err = <nil>

	fmt.Printf("w = %q\n", w) //	w = "abcdefghijklmnopqrstuvwxyz"

	// 尝试再次从 br 中读取数据
	b, err := br.ReadByte()
	fmt.Printf("b = %c, err = %v\n", b, err) //	b = , err = EOF
}

func aboutWrite() {
	/**
	实例化相关
	*/
	// 创建一个默认长度的 bufio.Writer，默认长度 4096
	//testNewWriter()

	// 创建一个自定义长度的 bufio.Writer
	//testNewWriterSize()

	// 将内容写入到缓存中
	//testWrite()

	// 功能同 Write，只不过以字符串形式写入
	//testWriteString()

	// 写入单个字节
	//testWriteByte()

	// 写入单个 rune，支持中文
	//testWriteRune()

	// 从一个 io.Writer 的缓存中读取数据并写入到当前 Writer 中
	//testReadFrom()
}
