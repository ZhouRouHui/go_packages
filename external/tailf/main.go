package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "./my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,                                 // 当日志文件被切割移走等操作时，可以定位到新的日志文件并打开
		Follow:    true,                                 // 文件被挪走或者更名，可以以及追踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从某个位置开始读，在极端情况下，比如进程异常退出时，可以定位到之前最后读取的位置
		MustExist: false,                                // 文件是否必须存在才监听，为 false 时文件不存在也监听，当文件生成后可以获取到日志
		Poll:      true,                                 // 轮循监听
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line // 一行日志
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
