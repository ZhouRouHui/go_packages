package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// NewReaderSize 将一个 io.Reader 封装成带缓存的 bufio.Reader 对象并返回
// 缓存大小有 size 决定（如果小于 16，则会被设置成 16）
func testNewReaderSize() *bufio.Reader {
	newReaderSizeObj := bufio.NewReaderSize(r, 8192)
	//fmt.Printf("newReaderSizeObj = %v\n", newReaderSizeObj)
	return newReaderSizeObj
}

// NewReader 返回一个默认大小的带缓存的 bufio.Reader 对象，默认大小为 4096
func testNewReader() *bufio.Reader {
	newReaderObj := bufio.NewReader(r)
	//fmt.Printf("newReaderObj = %v\n", newReaderObj)
	return newReaderObj
}

// Peek 将缓存内容中前 n 个字节的数据作为一个切片返回。该操作不会将数据读出，只是引用，
// 引用的数据在下一次读取操作之前是有效的。如果切片长度小于 n，则返回一个错误信息说明。
// 如果 n 大于缓存的总大小，则返回 ErrBufferFull
func testPeek() {
	peekRes, err := br.Peek(5)
	fmt.Printf("peekRes = %s, err = %v\n", peekRes, err) // abced

	// 改变缓冲区的内容，会使得原来的缓冲区里的内容也改变
	peekRes[0] = '#'
	peekRes, err = br.Peek(5)
	fmt.Printf("peekRes = %s, err = %v\n", peekRes, err) // #bcde
}

// 从 bufio.Reader 中读取数据到 p 中，返回写入 p 的字节数
// 读取到达结尾时，返回值 n 将为 0，而 err 将为 io.EOF
// 如果缓存不为空，则只能读出缓存中的数据，不从底层 io.Reader 中读取数据，
// 如果缓存为空，则：
// 		1. len(p) >= 缓存大小，则跳过缓存，直接从底层 io.Reader 中读出到 p 中。
// 		2. len(p) < 缓存大小，则先将数据从侧层 io.Reader 中读取到缓存中，再从缓存读取到 p 中。
func testRead() {
	b := make([]byte, 10)
	n, err := br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // abcdefghij 10 <nil>

	n, err = br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // klmnopqrst 10 <nil>

	n, err = br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // uvwxyz 6 <nil>

	n, err = br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // 0 EOF
}

// 返回可以从缓存中读取的字节数长度
func testBuffered() {
	b := make([]byte, 10)
	// 第一次读取
	br.Read(b)
	bufferedRes := br.Buffered()
	fmt.Printf("第一次读取后缓存中还有 %d 个长度可读\n", bufferedRes)

	br.Read(b)
	bufferedRes = br.Buffered()
	fmt.Printf("第二次读取后缓存中还有 %d 个长度可读\n", bufferedRes)
}

// 读取并返回一个字节，如果没有可用的数据，会返回错误
func testReadByte() {
	b, err := br.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%c\n", b)

	// 由于上面已经读出过一个字节，所以循环到 25 的时候会报错 EOF
	for i := 0; i < 26; i++ {
		b, err = br.ReadByte()
		if err != nil {
			panic(err)
		}
	}
}

// 撤销最后读出的字节
func testUnreadByte() {
	b, _ := br.ReadByte()
	fmt.Printf("%q\n", b)
	n := br.Buffered()
	fmt.Printf("剩余可读数量 %d\n", n)

	br.UnreadByte()
	n = br.Buffered()
	fmt.Printf("剩余可读数量 %d\n", n)
}

// 从缓存中读取一个 rune，支持中文
func testReadRune() {
	rr := strings.NewReader("中华人名共和国万岁")
	runeReader := bufio.NewReader(rr)
	tmp, size, _ := runeReader.ReadRune()
	fmt.Printf("读出的 tmp = %q, 长度 size = %d\n", tmp, size)
	n := runeReader.Buffered()
	fmt.Printf("剩余可读数量 %d\n", n) // 24 一个中文三个字节
}

func testUnreadRune() {
	rr := strings.NewReader("中华人名共和国万岁")
	runeReader := bufio.NewReader(rr)
	tmp, size, _ := runeReader.ReadRune()
	fmt.Printf("读出的 tmp = %q, 长度 size = %d\n", tmp, size)
	n := runeReader.Buffered()
	fmt.Printf("剩余可读数量 %d\n", n) // 24 一个中文三个字节

	runeReader.UnreadRune()
	n = runeReader.Buffered()
	fmt.Printf("剩余可读数量 %d\n", n) // 27 一个中文三个字节
}

// ReadLine 是一个低水平的行读取原语，大多数情况下，应该使用 ReadBytes('\n') 或 ReadString('\n'), 或者使用一个 Scanner。
//
// ReadLine 通过调用 ReadSlice 方法实现，返回的也是缓存的切片。用于读取一行数据，不包括行尾标记(\n 或 \r\n)
//
// 只要能读出数据，err 就为 nil。如果没有数据可读，则 isPrefix 返回 false，err 返回 io.EOF.
//
// 如果找到行尾标记，则返回查找结果，isPrefix 返回 false。
// 如果未找到行尾标记，则：
// 		1. 缓存不满，则将缓存填满后再次查找
// 		2. 缓存是满的，则返回整个缓存，isPrefix 返回 true
//
// 整个数据尾部 `有一个换行标记` 和 `没有换行标记` 的读取结果是一样的。
//
// 如果 ReadLine 读取到换行标记，则调用 UnreadByte 撤销的是换行标记，而不是返回的数据
func testReadLine() {
	readLine := strings.NewReader("abc\ndef\nghi")
	readLineReader := bufio.NewReader(readLine)
	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = readLineReader.ReadLine()
		fmt.Printf("%q %t %v\n", line, isPrefix, err)
	}
}

// ReadSlice 在 b 中查找 delim 并发挥 delim 及其之前的所有数据。
// 该操作会读出数据，返回的切片是已读出的数据的引用，切片中的数据在下一次读取操作之前是有效的
//
// 如果找到 delim，则返回查找结果，err 返回 nil
// 如果未找到 delim，则：
// 		1、缓存不满，则将缓存填满后再次查找。
// 		2、缓存是满的，则返回整个缓存，err 返回 ErrBufferFull。
//
// 如果未找到 delim 且遇到错误（通常是 io.EOF），则返回缓存中的所有数据和遇到的错误。
//
// 因为返回的数据有可能被下一次的读写操作修改，所以大多数操作应该使用 ReadBytes 或 ReadString，
// 它们返回的是数据的拷贝。
func testReadSlice() {
	sr := strings.NewReader("abc def ghi")
	readSliceReader := bufio.NewReader(sr)

	b, err := readSliceReader.ReadSlice(' ')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)

	b, err = readSliceReader.ReadSlice(' ')
	if err != nil {
		panic(err)
	}

	fmt.Printf("%q\n", b)
	b, err = readSliceReader.ReadSlice(' ')
	if err != nil {
		// 第三次读取找不到 delim 会报错
		panic(err)
	}
	fmt.Printf("%q\n", b)
}

// 功能同 ReadSlice，只不过返回的是缓存的拷贝。
func testReadBytes() {
	s := strings.NewReader("ABC,EFG,HIJ")
	readBytesReader := bufio.NewReader(s)
	b, err := readBytesReader.ReadBytes(',')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)
	b, err = readBytesReader.ReadBytes(',')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)
	b, err = readBytesReader.ReadBytes(',')
	if err != nil {
		// 第三次读取找不到 delim 会报错
		panic(err)
	}
	fmt.Printf("%q\n", b)
}

// 功能同 ReadString，但是返回的是字符串
func testReadString() {
	s := strings.NewReader("ABC,EFG,HIJ")
	readBytesReader := bufio.NewReader(s)
	b, err := readBytesReader.ReadString(',')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)
	b, err = readBytesReader.ReadString(',')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)
	b, err = readBytesReader.ReadString(',')
	if err != nil {
		// 第三次读取找不到 delim 会报错
		panic(err)
	}
	fmt.Printf("%q\n", b)
}

// 将缓存中的内容读出并写入到某地
// WriteTo 方法实现了 io.WriterTo 接口
func testWriteTo() {
	b := bytes.NewBuffer(make([]byte, 0))
	br.WriteTo(b)
	fmt.Printf("%q\n", b) // "abcdefghijklmnopqrstuvwxyz"

	// 尝试再从缓存中读取内容
	rb := make([]byte, 10)
	n, err := br.Read(rb)
	fmt.Printf("rb = %s, n = %d, err=%v\n", string(rb), n, err) // rb = , n = 0, err=EOF 因为缓存的内容在 WriteTo 的时候已经读完，这里已经读不到内容
}

func aboutReader() {
	/**
	实例化相关
	*/
	// 创建一个自定义长度的 bufio.Reader
	//_ = testNewReaderSize()

	// 创建一个默认长度的 bufio.Reader 默认长度为 4096
	//_ = testNewReader()

	/**
	读操作
	*/
	// 引用但不从缓存中读出
	//testPeek()

	// 从缓存中读出内容
	//testRead()

	// 返回可以从缓存中读取的字节数长度
	//testBuffered()

	// 从缓存中读取一个字节
	//testReadByte()

	// 撤销最后读出的字节
	//testUnreadByte()

	// 从缓存中读取一个 rune，支持中文
	//testReadRune()

	// 撤销最后读出的 rune, 支持中文
	//testUnreadRune()

	// 读取一行数据
	//testReadLine()

	// 从缓存中寻找 delim，并返回包括 delim 和之前的内容，是数据不安全的，不建议使用
	//testReadSlice()

	// 功能同 ReadSlice，是数据安全的
	//testReadBytes()

	// 功能同 ReadBytes，但返回的是字符串
	//testReadString()

	// bufio.Reader 下的写
	// 将缓存中的内容读出并写入到某地
	//testWriteTo()
}
