package main

import (
	"fmt"
	"path"
)

/**
path 包

概述
path 包实现了一系列使用函数，这些函数可以对可以对用斜杠进行分割的路径进行处理。如果你要对操作系统路径进行处理，请使用 path/filepath 包。

变量
ErrBadPattern 用于指示一个 glob 模式出现了模式格式。

Base 函数
func Base(path string) string
返回路径的最后一个元素。在提取最后一个元素之前，路径末尾的斜杠会被移除。如果路径为空，那么返回"."。如果路径只有斜杠组成，那么返回"/"。
`
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b"))	// b
}
`

Clean 函数
func Clean(path string) string
通过纯粹的词法处理，返回一个等价于 path 的最短路径。Clean 函数会根据以下规则进行迭代，直到没有任何规则可以应用为止：
	1. 使用单个斜杠替代多个斜杠
	2. 移除所有 . 路径名元素（当前目录）
	3. 移除所有内涵的 .. 路径名元素（父目录），以及先于他的非 .. 元素
	4. 移除所有以根路径方式存在的 .. 元素，也即是，使用 "/" 去替换路径开头的 "/.."
只有在 path 为根目录 "/" 的时候，函数才会返回 "/"。
如果函数的处理结果是一个空字符串，那么他将返回字符串 "." 作为结果
示例：
`
package main

import (
	"fmt"
	"path"
)

func main() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}
}

执行结果：
Clean("a/c") = "a/c"
Clean("a//c") = "a/c"
Clean("a/c/.") = "a/c"
Clean("a/c/b/..") = "a/c"
Clean("/../a/c") = "/a/c"
Clean("/../a/b/../././/c") = "/a/c"
`

Dir 函数
func Dir(path string) string
返回路径里面除最后一个元素之外的其他所有元素，通常为路径的目录。
在使用 Split 丢弃最后一个元素之后，路径将是清洁的（Cleaned），并且所有末尾的斜杠都将被移除。
如果路径为空，那么函数将返回 "." 作为结果。
如果路径由任意多个斜杠后跟任意多个非斜杠字节组成（比如 //////qqq），那么函数将返回单个斜杠。在其他情况下，函数返回的路径都不会以斜杠结束。
示例：
`
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("/a/b/c"))
}
`

Ext 函数
func Ext(path string) string
返回路径使用的文件扩展名。
示例：
`
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Ext("/a/b/c/bar.css"))	// .css
}
`

IsAbs 函数
func IsAbs(path string) bool
检查路径是否为绝对路径
示例：
`
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.IsAbs("/dev/null"))
}
`

Join 函数
func Join(elem ...string) string
将任意数量的路径元素拼接成为单个路径，并在有需要时添加用于分隔的斜杠。函数返回的结果路径是 Cleaned 的，所有空字符串都将被忽略。
示例：
`
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Join("a/b", "/c"))
}
`

Match 函数
func Match(pattern, name string) (matched bool, err error)
返回文件名是否与给定规则配。规则语法：
	模式:
		{ 匹配项 }

	匹配项:
		'*'         匹配任意长度的非 / 字符
		'?'         匹配单个非 / 字符
		'[' [ '^' ] { 字符范围 } ']'
					字符类型必须是非空的
		c           匹配不等于 '*' 、 '?' 、 '\\' 和 '[' 的字符 c
		'\\' c      匹配字符 c

	字符范围
		c           匹配不等于 '\\' 、 '-' 、 和 ']' 的字符 c
		'\\' c      匹配字符 c
		lo '-' hi   匹配字符 c ，其中 lo <= c <= hi
Match 要求模式与文件名完全匹配，而不是部分匹配。函数唯一可能返回的错误为 ErrBadPattern，这个错误用于表示给定的模式有误。

Split 函数
func Split(path string) (dir, file string)
根据路径的最有一个斜杠，将路径划分为目录部分和文件名部分。
如果路径中不包含斜杠，那么返回一个空的目录，并将文件名部分设置为路径本身。返回值满足这样一个属性，也即是 path = dir + file.
示例：
`
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Split("static/myfile.css"))	// static/ myfile.css
}
`
*/

func testBase() {
	fmt.Printf("testBase res: %s\n", path.Base("/a/b")) // b
	fmt.Println()
}

func testClean() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
	}
	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}
	fmt.Println()
}

func testDir() {
	fmt.Printf("testDir res: %s\n", path.Dir("/a/b/c")) // /a/b
	fmt.Println()
}

func testExt() {
	fmt.Printf("testExt res: %s\n", path.Ext("/a/b/c.css")) // .css
	fmt.Println()
}

func testIsAbs() {
	fmt.Printf("testIsAbs res: %t\n", path.IsAbs("/dev/null")) // true
	fmt.Println()
}

func testJoin() {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Join("a/b", "/c"))
	fmt.Println()
}

func testSplit() {
	fmt.Println(path.Split("static/myfile.css")) // static/ myfile.css
}

func main() {
	testBase()

	testClean()

	testDir()

	testExt()

	testIsAbs()

	testJoin()

	testSplit()
}
