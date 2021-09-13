package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var n int = 999

func globalFunc() {
	fmt.Println("全局函数 ------------------------- start --------------------------")
	// New() 方法测试
	newFuncTest()

	// 获取一个默认的系统定义的 Logger
	defaultRes := log.Default()
	fmt.Printf("defaultRes = %v\n", defaultRes)

	/**
	Print 系列
	*/
	log.Println("log.Println func test info")
	log.Printf("log.Printf func test info, %d\n", n)
	log.Print("log.Print func test info")

	/**
	Fatal 系列，输出日志后，程序中断
	*/
	//log.Fatalln("log.Fatalln func test info")
	//log.Fatalf("log.Fatalf func test into, %d\n", n)
	//log.Fatal("log.Fatal func test info")

	/**
	Panic 系列，输出日志后，执行 panic
	*/
	defer panicHandle()
	//log.Panicln("log.Panicln func test info")
	//log.Panicf("log.Panicf func test info, %d\n", n)
	//log.Panic("log.Panic func test info")

	// 返回日志输出位置
	writeRes := log.Writer()
	fmt.Printf("log 包默认的输出位置 = %v\n", writeRes)

	// 设置日志输出位置, 系统默认的事 os.Stderr
	log.SetOutput(os.Stdout)
	writeRes = log.Writer()
	fmt.Printf("writeRes = %v\n", writeRes)

	// 返回日志记录器的 flag 值
	flagsRes := log.Flags()
	fmt.Printf("flagsRes = %d\n", flagsRes)

	// 设置日志记录器的 flag 值
	log.SetFlags(19)
	flagsRes = log.Flags()
	fmt.Printf("flagsRes = %d\n", flagsRes)
	log.SetFlags(3) // 将 flag 值重置回去

	// 返回日志记录器的 prefix
	prefixRes := log.Prefix()
	fmt.Printf("prefixRes = %s\n", prefixRes)

	// 设置日志记录器的 prefix
	log.SetPrefix("Test Logger from default log: ")
	prefixRes = log.Prefix()
	fmt.Printf("prefixRes = %s\n", prefixRes)

	fmt.Println("全局函数 ------------------------- end --------------------------")
}

// newFuncTest New() 方法测试
func newFuncTest() {
	// New() 函数创建一个新的定制化的 Logger
	// 参数一设置日志将被写入的目的地
	// 参数二会在生成的每行日志的最开始出现
	// 参数三定义日志记录包含哪些属性
	file, err := os.OpenFile("./errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open error log file: ", err)
	}
	var (
		Trace   *log.Logger //记录所有日志
		Info    *log.Logger //重要的信息
		Warning *log.Logger //需要注意的信息
		Error   *log.Logger //非常严重的问题
	)
	// 创建定制化的 Logger
	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	// 往 Logger 中写入日志
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

// loggerFunc Logger 结构体方法
func loggerFunc() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Logger 结构体方法 ------------------------- start --------------------------")
	fmt.Println()

	// 为了测试方便生成一个终端输出的日志记录器
	stdoutLogger := createStdoutLogger()

	// Output 输出日志，这是底层的输出日志方法，其他的 Print 等都是调用这个方法
	// todo 没懂为什么参数一都写成 2 ？
	stdoutLogger.Output(2, "log.Logger.Output func test info")

	/**
	Print 系列输出日志
	*/
	stdoutLogger.Printf("log.Logger.Printf func test info, %d", n)
	stdoutLogger.Println("log.Logger.Println func test info")
	stdoutLogger.Print("log.Logger.Print func test info")

	/**
	Fatal 系列输出日志，会将程序终止
	*/
	//stdoutLogger.Fatalf("log.Logger.Fatalf func test info, %d", n)
	//stdoutLogger.Fatalln("log.Logger.Fatalln func test info")
	//stdoutLogger.Fatal("log.Logger.Fatal func test info")

	/**
	Panic 系列，在日志被输出后会执行 panic
	*/
	defer panicHandle()
	//stdoutLogger.Panicf("log.Logger.Panicf func test info, %d", n)
	//stdoutLogger.Panicln("log.Logger.Panicln func test info")
	//stdoutLogger.Panic("log.Logger.Panic func test info")

	// 返回 Logger 的 flag 值
	flagsRes := stdoutLogger.Flags()
	fmt.Printf("flagsRes = %d\n", flagsRes)

	// 设置 Logger 的 flag 值
	stdoutLogger.SetFlags(3)
	flagsRes = stdoutLogger.Flags()
	fmt.Printf("flagsRes = %d\n", flagsRes)
	stdoutLogger.SetFlags(19) // 将其重置为 19

	// 返回 Logger 的 prefix 内容
	prefixRes := stdoutLogger.Prefix()
	fmt.Printf("prefixRes = %s\n", prefixRes)

	// 设置 Logger 的 prefix 值
	stdoutLogger.SetPrefix("Test Logger from SetPrefix: ")
	prefixRes = stdoutLogger.Prefix()
	fmt.Printf("prefixRes = %s\n", prefixRes)

	// 返回 Logger 的输出目的地
	writeRes := stdoutLogger.Writer()
	fmt.Printf("writeRes = %v\n", writeRes)

	// 设置 Logger 输出目的地
	testSetOutput()

	fmt.Println("Logger 结构体方法 ------------------------- end --------------------------")

}

// testSetOutput 测试 log.Logger.SetOutput() 方法
func testSetOutput() {
	// 系统默认的 Logger 的 out 是 os.Stderr，把他改成 os.Stdout 后，可以看到前后两次打印的 out 的内存地址不一样了
	setOutputObj := log.Default()
	fmt.Printf("testSetOutput: %+v\n", setOutputObj)
	setOutputObj.SetOutput(os.Stdout)
	fmt.Printf("testSetOutput: %+v\n", setOutputObj)
}

// createStdoutLogger 创建一个 os.Stdout 的日志记录器
func createStdoutLogger() *log.Logger {
	return log.New(os.Stdout, "Test Logger: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func panicHandle() {
	if err := recover(); err != nil {
		fmt.Printf("this is from recover func, err = %v\n", err)
	}
}

func main() {
	/**
	全局函数
	*/
	globalFunc()

	/**
	Logger 结构体方法
	*/
	loggerFunc()
}
