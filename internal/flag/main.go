package main

import (
	"flag"
	"fmt"
	"time"
)

/**
flag 包设置命令行参数
简易教程: https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/flag.html
 */
func main() {
	/**
	命令行参数分为两种
	1. flag, 比如 go run main.go -env=test, 其中 -env=test 就是一个 flag，flag 需要在程序中定义
	2. arg, 比如 go run main.go a b c, 其中 a，b，c 就是 arg，arg 不需要在程序中定义，只需要在程序中直接获取就可以

	定义 flag 有两种方式，其中的 Type 表示（Int, String... 等类型）
	方式一: flag.Type("参数名", "默认值", "使用帮助")	返回的内容是对应参数的指针
	方式二: flag.TypeVar("变量指针", "参数名", "默认值", "使用帮助")	将参数值写入到第一个参数变量指针里面
	两种方式要想命令行参数有效写入到对应变量，都需要执行 flag.Parse() 方法
	 */


	/**
	flag 相关操作
	 */
	fmt.Println("flag 相关操作 --------------------- start ----------------------------------")
	//name := flag.String("name", "zrh", "姓名")
	//age := flag.Int("age", 18, "年龄")
	//married := flag.Bool("married", false, "婚否")
	//delay := flag.Duration("d", 0, "时间间隔")
	var (
		name string
		age int
		married bool
		delay time.Duration
	)
	// string 类型
	flag.StringVar(&name, "name", "zrh", "-name=zrh")
	// int 类型
	flag.IntVar(&age, "age", 18, "-age=18")
	// bool 类型
	flag.BoolVar(&married, "married", false, "-married=false")
	// duration 类型
	flag.DurationVar(&delay, "delay", 100, "-delay=30m0s")

	// 把用户传递的命令行参数解析为对应变量的值，必须使用，否则拿不到命令行的参数
	flag.Parse()

	parsedRes := flag.Parsed()
	fmt.Printf("parsedRes = %t\n", parsedRes)

	fmt.Println(name, age, married, delay)

	// 对命令行参数赋值
	_ = flag.Set("name", "loedan")
	fmt.Println(name)	// loedan

	// 打印所有命令行标志的用法和帮助信息
	flag.PrintDefaults()

	// 命令行设置的参数个数
	nFlagRes := flag.NFlag()
	fmt.Printf("nFlagRes = %d\n", nFlagRes)

	fmt.Println("flag 相关操作 --------------------- end ----------------------------------")



	/**
	arg 相关操作
	 */
	fmt.Println()
	fmt.Println()
	fmt.Println("arg 相关操作 --------------------- start ----------------------------------")

	// 根据下标获取 arg，如果不存在则返回空字符串
	argRes := flag.Arg(1)
	fmt.Printf("argRes = %s\n", argRes)

	// 返回 arg 的总个数
	nArgRes := flag.NArg()
	fmt.Printf("nArgRes = %d\n", nArgRes)


	argsRes := flag.Args()
	fmt.Printf("argsRes = %v\n", argsRes)
	if len(flag.Args()) > 0 {
		for i, v := range flag.Args() {
			fmt.Printf("flag.Args[%d]=%v\n", i, v)
		}
	}
}
