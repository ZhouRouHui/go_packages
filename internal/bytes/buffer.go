package main

import (
	"bytes"
	"fmt"
)

/**
// 将 buf 包装成 bytes.Buffer 对象。
func NewBuffer(buf []byte) *Buffer

// 将 s 转换为 []byte 后，包装成 bytes.Buffer 对象。
func NewBufferString(s string) *Buffer

// Buffer 本身就是一个缓存（内存块），没有底层数据，缓存的容量会根据需要
// 自动调整。大多数情况下，使用 new(Buffer) 就足以初始化一个 Buffer 了。

// bytes.Buffer 实现了如下接口：
// io.ReadWriter
// io.ReaderFrom
// io.WriterTo
// io.ByteWeriter
// io.ByteScanner
// io.RuneScanner

// 未读取部分的数据长度
func (b *Buffer) Len() int

// 缓存的容量
func (b *Buffer) Cap() int

// 读取前 n 字节的数据并以切片形式返回，如果数据长度小于 n，则全部读取。
// 切片只在下一次读写操作前合法。
func (b *Buffer) Next(n int) []byte

// 读取第一个 delim 及其之前的内容，返回遇到的错误（一般是 io.EOF）。
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
func (b *Buffer) ReadString(delim byte) (line string, err error)

// 写入 r 的 UTF-8 编码，返回写入的字节数和 nil。
// 保留 err 是为了匹配 bufio.Writer 的 WriteRune 方法。
func (b *Buffer) WriteRune(r rune) (n int, err error)

// 写入 s，返回写入的字节数和 nil。
func (b *Buffer) WriteString(s string) (n int, err error)

// 引用未读取部分的数据切片（不移动读取位置）
func (b *Buffer) Bytes() []byte

// 返回未读取部分的数据字符串（不移动读取位置）
func (b *Buffer) String() string

// 自动增加缓存容量，以保证有 n 字节的剩余空间。
// 如果 n 小于 0 或无法增加容量则会 panic。
func (b *Buffer) Grow(n int)

// 将数据长度截短到 n 字节，如果 n 小于 0 或大于 Cap 则 panic。
func (b *Buffer) Truncate(n int)

// 重设缓冲区，清空所有数据（包括初始内容）。
func (b *Buffer) Reset()
*/

func aboutBytesBuffer() {
	rd := bytes.NewBufferString("Hello World!")
	buf := make([]byte, 6)
	// 获取数据切片
	b := rd.Bytes()
	// 读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd.String()) // World!
	fmt.Printf("%s\n\n", b)         // Hello World!

	// 写入一部分数据，看看切片有没有变化
	rd.Write([]byte("abcdefg"))
	fmt.Printf("%s\n", rd.String()) // World!abcdefg
	fmt.Printf("%s\n\n", b)         // Hello World!

	// 再读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd.String()) // abcdefg
	fmt.Printf("%s\n", b)           // Hello World!
}
