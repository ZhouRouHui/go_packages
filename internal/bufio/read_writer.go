package main

import (
	"bufio"
	"fmt"
)

// ReaderWriter 类型保管了指向 Reader 和 Writer 类型的指针，实现了 io.ReadWriter 接口
func testReadWriter() {
	rw := bufio.NewReadWriter(br, bw)

	// 读取数据
	str, err := rw.ReadString('g')
	fmt.Printf("str = %s, err = %v\n", str, err) // str = abcdefg, err = <nil>

	// 写入数据
	n, err := rw.WriteString("abcdefg")
	fmt.Printf("n = %d, err = %v\n", n, err) // n = 7, err = <nil>

	rw.Flush()
	fmt.Printf("w = %q\n", w) // abcdefg
}
