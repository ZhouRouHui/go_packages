package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

/**
os/exec 包提供一个渠道可以在程序中执行外部命令
*/

func main() {
	// 在环境变量 path 中查找可执行二进制文件
	// 返回完整路径或者相对于当前目录的相对路径
	lookPathRes, _ := exec.LookPath("go")
	fmt.Printf("lookPathRes = %s\n", lookPathRes)

	// 创建并返回一个 exec.Cmd
	cmd := exec.Command("go", "version")
	// 执行命令，并返回标准输出和错误输出
	combineOutputRes, _ := cmd.CombinedOutput()
	fmt.Printf("combineOutputRes = %s\n", string(combineOutputRes))

	// 创建一个 exec.Cmd
	cmd2 := exec.Command("ls", "-al")
	buf := bytes.Buffer{}
	// 将 cmd2 的标准输出地址设置为 buf 的地址，也就是让 cmd2 的标准输出的内容写在 buf 的内存地址中
	cmd2.Stdout = &buf
	// 执行命令，阻塞直到完成
	cmd2.Run()
	fmt.Printf("cmd2.Stdout = %s\n", buf.String())

	// 创建一个 exec.Cmd
	cmd3 := exec.Command("git", "status")
	// 获取命令在 start 后标准输出管道
	out3, _ := cmd3.StdoutPipe()
	// 执行命令
	cmd3.Start()
	// 读取管道中所有数据
	data3, _ := ioutil.ReadAll(out3)
	// 等待命令执行完成
	cmd3.Wait()
	fmt.Printf("cmd3 的管道数据 = %s\n", string(data3))
}
