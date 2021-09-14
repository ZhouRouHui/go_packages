package main

// https://www.cnblogs.com/saryli/p/11691142.html
/**
os 系统包详解
*/

/**
环境变量相关
func Environ() []string 获取所有环境变量，返回变量列表
`
	package main

	import (
		"fmt"
		"os"
		"strings"
	)

	func main() {
		envs := os.Environ()
		for _, env := range envs {
			cache := strings.Split(env, "=")
			fmt.Printf(`
			key: %s value: %s
			`, cache[0], cache[1])
		}
	}
`

func Getenv(key string) string	获取指定环境变量
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		fmt.Println(os.Getenv("GOPATH"))
	}
`

func Setenv(key, value string) error 设置环境变量
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		fmt.Println(os.Getenv("GOPATH"))

		if err := os.Setenv("GOPATH", "./GO/bin"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	}
`

func Clearenv()	清除所有环境变量，慎用
`
	os.Clearenv()
`
*/

/**
文件模式
const (
    // 单字符是被String方法用于格式化的属性缩写。
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
    ModeAppend                                     // a: 只能写入，且只能写入到末尾
    ModeExclusive                                  // l: 用于执行
    ModeTemporary                                  // T: 临时文件（非备份文件）
    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    ModeDevice                                     // D: 设备
    ModeNamedPipe                                  // p: 命名管道（FIFO）
    ModeSocket                                     // S: Unix域socket
    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
)
*/

/**
文件信息相关
FileInfo 接口
`
	type FileInfo interface {
		Name() string       // 文件的名字（不含扩展名）
		Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		Mode() FileMode     // 文件的模式位
		ModTime() time.Time // 文件的修改时间
		IsDir() bool        // 等价于Mode().IsDir()
		Sys() interface{}   // 底层数据来源（可以返回nil）
	}
`

func Stat(name string) (fi FileInfo, err error)	获取文件信息对象, 符号链接将跳转
`
	fi, _ := os.Stat("./cache.js")
	fmt.Println(fi.Size())
`

func Lstat(name string) (FileInfo, error)	获取文件信息对象, 符号链接不跳转
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		fi, _ := os.Lstat("./main.go")
		fmt.Println(fi.Size())
	}
`

func IsExist(err error) bool 根据错误，判断 文件或目录是否存在
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		if _, err := os.Open("./empty.js"); err != nil {
			// false 不存在   true 存在
			emptyErr := os.IsExist(err)
			fmt.Println(emptyErr, "\n", err)
		}
	}
`

func IsNotExist(err error) bool	IsExist 反义方法
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		if _, err := os.Open("./empty.js"); err != nil {
			// false 不存在   true 存在
			emptyErr := os.IsNotExist(err)
			fmt.Println(emptyErr, "\n", err)
		}
	}
`

func IsPermission(err error) bool	根据错误，判断是否为权限错误
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		file, _ := os.Open("cache.js")
		_, err := file.WriteString("// new info")

		if err != nil {
			fmt.Println(os.IsPermission(err))
		}
		defer file.Close()
	}
`
*/

/**
文件目录操作

属性操作
func Getwd() (dir string, err error)	获取当前工作目录
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		path, _ := os.Getwd()
		fmt.Println(path)
	}
`

func Chdir(dir string) error	修改当前，工作目录
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		path1, _ := os.Getwd()
		fmt.Println(path1)
		os.Chdir("./../")
		path, _ := os.Getwd()
		fmt.Println(path)

	}
`

func Chmod(name string, mode FileMode) error	修改文件的 FileMode

func Chtimes(name string, atime time.Time, mtime time.Time) error	修改文件的 访问时间和修改时间
`
	package main

	import (
		"fmt"
		"os"
		"time"
	)

	func main() {
		fmt.Println(os.Getwd())

		path := "test.txt"
		os.Chtimes(path, time.Now(), time.Now())

		fi, _ := os.Stat(path)
		fmt.Println(fi.ModTime())

	}
`

增删改查
func Mkdir(name string, perm FileMode) error	创建目录
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		if err := os.Mkdir("test", os.ModeDir); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	}
`

func MkdirAll(path string, perm FileMode) error	递归创建目录
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		if err := os.MkdirAll("test01/test", os.ModeDir); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	}
`

func Remove(name string) error	移除文件或目录(单一文件)
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		if err := os.Remove("test"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	}
`

func RemoveAll(path string) error	递归删除文件或目录
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		if err := os.RemoveAll("test01"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	}
`

func Rename(oldpath, newpath string) error	文件重名或移动
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {

		// 重命名
		err := os.Rename("test.txt", "test01.js")
		if err != nil {
			fmt.Println(err)
		}
		err = os.Mkdir("test", os.ModeDir)
		if err != nil {
			fmt.Println(err)
		}

		// 移动
		err = os.Rename("test01.js", "test/text01.txt")
		if err != nil {
			fmt.Println(err)
		}
	}
`

func Truncate(name string, size int64) error	修改文件大小
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {

		path := "test/text01.txt"
		fi, err := os.Stat(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		size := fi.Size()
		fmt.Println(size)

		// 截取长度
		size = int64(float64(size) * 0.5)

		os.Truncate(path, size)

		fi, _ = os.Stat(path)

		fmt.Println(fi.Size())
	}
`

func SameFile(fi1, fi2 FileInfo) bool	比较两个文件信息对象，是否指向同一文件
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {

		path := "test/text01.txt"

		fi_1, _ := os.Stat(path)
		fi_2, _ := os.Stat(path)

		fmt.Println(os.SameFile(fi_1, fi_2))
	}
`
*/

/**
文件目录对象

打开文件/目录
func Create(name string) (file *File, err error)	创建文件, 如果文件存在，清空原文件
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		file, _ := os.Create("./new_file.js")
		fmt.Println(file.Name())
	}
`

func Open(name string) (file *File, err error)	打开文件，获取文件对象, 以读取模式打开
Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。
如果出错，错误底层类型是*PathError。所以，Open()只能用于读取文件。
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		file, _ := os.Open("./new_file.js")
		fmt.Println(file.Name())
	}
`

func OpenFile(name string, flag int, perm FileMode) (file *File, err error)	以指定模式，打开文件
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {

		file, _ := os.OpenFile("./new_file.js", os.O_RDONLY, os.ModeAppend)
		fmt.Println(file.Name())
	}
`

文件对象属性操纵
func (f *File) Name() string	获取文件路径

func (f *File) Stat() (fi FileInfo, err error)	获取文件信息对象

func (f *File) Chdir() error	将当前工作路径修改为文件对象目录， 文件对象必须为目录, 该接口不支持window

func (f *File) Chmod(mode FileMode) error	修改文件模式

func (f *File) Truncate(size int64) error	修改文件对象size


文件对象读写操作
func (f *File) Read(b []byte) (n int, err error)	读取文件内容, 读入长度取决 容器切片长度
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		bt := make([]byte, 10)
		file, _ := os.Open("./new_file.js")

		file.Read(bt)
		defer file.Close()

		fmt.Println(string(bt))
	}
`

func (f *File) ReadAt(b []byte, off int64) (n int, err error)	从某位置，读取文件内容
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {

		bt := make([]byte, 100)
		file, _ := os.Open("test/text01.txt")

		file.ReadAt(bt, 2)

		fmt.Println(string(bt))
	}
`

func (f *File) Write(b []byte) (n int, err error)	写入内容
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {

		file, err := os.OpenFile("test/text01.txt", os.O_RDWR, os.ModeAppend)
		if err != nil {
			fmt.Println("err: ", err)
			os.Exit(1)
		}

		defer file.Close()

		if n, err := file.Write([]byte("// new info")); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n)
		}
	}
`

func (f *File) WriteString(s string) (ret int, err error)	写入字符
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		file, err := os.OpenFile("test/text01.txt", os.O_RDWR, os.ModeAppend)
		if err != nil {
			fmt.Println("err: ", err)
			os.Exit(1)
		}

		defer file.Close()

		if n, err := file.WriteString("// test info"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n)
		}
	}
`

func (f *File) WriteAt(b []byte, off int64) (n int, err error)	从指定位置，写入
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		file, err := os.OpenFile("test/text01.txt", os.O_RDWR, os.ModeAppend)
		if err != nil {
			fmt.Println("err: ", err)
			os.Exit(1)
		}

		defer file.Close()

		if n, err := file.WriteAt([]byte(" append "), 5); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n)
		}
	}
`

func (f *File) Seek(offset int64, whence int) (ret int64, err error)	设置下次读写位置
`
	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		f, err := os.OpenFile("test/text01.txt", os.O_RDWR, os.ModeAppend)
		if err != nil {
			fmt.Println("err: ", err)
			os.Exit(1)
		}

		defer f.Close()

		f.Seek(2, 0)
		buffer := make([]byte, 5)
		// Read 后文件指针也会偏移

		n, err := f.Read(buffer)
		if err != nil {
			fmt.Println(nil)
			return
		}
		fmt.Printf("n is %d, buffer content is : %s\n", n, buffer)
		// 获取文件指针当前位置
		cur_offset, _ := f.Seek(0, os.SEEK_CUR)
		fmt.Printf("current offset is %d\n", cur_offset)
	}
`

func (f *File) Close() error	关闭文件
*/

func main() {
	
}
