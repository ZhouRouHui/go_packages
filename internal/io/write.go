package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFuncTest() {
	// string 类型写入，为什么需要这个？
	// io.Write 接口接收的是 []byte 类型，所以当我们需要写入内容的时候都需要显示的进行类型转换 []byte("something to write")，这意味着要开辟一块临时内存，
	// 如果调用 write 的次数太多，开辟的内存次数太多了，会影响应用的性能，所以提供了这样一个方法。
	// 方法内部会进行类型断言，如果是 StringWriter 接口类型则直接 string 类型写入，否则再调用 Write() 方法并显示转换类型写入
	writeStringRes, _ := io.WriteString(os.Stdout, "io.WriteString func test info")
	fmt.Printf("writeStringRes = %d\n", writeStringRes)
}
