package main

// https://www.cnblogs.com/golove/p/5903579.html
/**
path/filepath 包中的函数会根据不同平台做不同的处理，比如路径分隔符，卷名等。

路径分隔符转换：
`
const (
    Separator     = os.PathSeparator      // 路径分隔符（分隔路径元素）
    ListSeparator = os.PathListSeparator  // 路径列表分隔符（分隔多个路径）
)
`
下面两个函数主要用于将 Windows 路径分隔符转换为 Linux 路径分隔符，处理完后再转换回去，只在 Windows 中有用，在 Linux 中没必要：
// 将 path 中平台相关的路径分隔符转换为 '/'
ToSlash(path string) string
// 将 path 中的 '/' 转换为系统相关的路径分隔符
FromSlash(path string) string
`
func main() {
	s := `http://www.site.com/a/b/c/d`
	u, _ := url.Parse(s)
	s = u.Path
	// 下面这句用于 Windows 系统
	s = filepath.FromSlash(s)
	fmt.Println(s) // /a/b/c/d 或 \a\b\c\d
	// 创建目录试试
	if err := os.MkdirAll(s[1:], 0777); err != nil {
		fmt.Println(err)
	}
	// 下面这句用于 Windows 系统
	s = filepath.ToSlash(s)
	fmt.Println(s) // /a/b/c/d
}
`



// 获取 path 中最后一个分隔符之前的部分（不包含分隔符）
Dir(path string) string
// 获取 path 中最后一个分隔符之后的部分（不包含分隔符）
Base(path string) string
// 获取 path 中最后一个分隔符前后的两部分
// 之前包含分隔符，之后不包含分隔符
Split(path string) (dir, file string)
// 获取路径字符串中的文件扩展名
Ext(path string) string
`
func main() {
	path := `a///b///c///d`
	path = filepath.FromSlash(path) // 平台处理

	d1 := filepath.Dir(path)
	f1 := filepath.Base(path)
	d2, f2 := filepath.Split(path)
	fmt.Printf("%q  %q\n%q  %q\n", d1, f1, d2, f2)
	// "a/b/c"  "d"
	// "a///b///c///"  "d"

	ext := filepath.Ext(path)
	fmt.Println(ext) // .txt
}
`



// 获取 targpath 相对于 basepath 的路径。
// 要求 targpath 和 basepath 必须“都是相对路径”或“都是绝对路径”。
Rel(basepath, targpath string) (string, error)
`
func main() {
	// 都是绝对路径
	s, err := filepath.Rel(`/a/b/c`, `/a/b/c/d/e`)
	fmt.Println(s, err) // d/e <nil>

	// 都是相对路径
	s, err = filepath.Rel(`a/b/c`, `a/b/c/d/e`)
	fmt.Println(s, err) // d/e <nil>

	// 一个绝对一个相对
	s, err = filepath.Rel(`/a/b/c`, `a/b/c/d/e`)
	fmt.Println(s, err)
	//  Rel: can't make a/b/c/d/e relative to /a/b/c

	// 一个相对一个绝对
	s, err = filepath.Rel(`a/b/c`, `/a/b/c/d/e`)
	fmt.Println(s, err)
	//  Rel: can't make /a/b/c/d/e relative to a/b/c

	// 从 `a/b/c` 进入 `a/b/d/e`，只需要进入 `../d/e` 即可
	s, err = filepath.Rel(`a/b/c`, `a/b/d/e`)
	fmt.Println(s, err) // ../d/e <nil>
}
`



// 将 elem 中的多个元素合并为一个路径，忽略空元素，清理多余字符。
Join(elem ...string) string
`
func main() {
	// Linux 示例
	s := filepath.Join("a", "b", "", ":::", "  ", `//c////d///`)
	fmt.Println(s) // a/b/:::/  /c/d
}
`



// 清理路径中的多余字符，比如 /// 或 ../ 或 ./
Clean(path string) string
`
func main() {
	// Linux 示例
	s := filepath.Clean("a/./b/:::/..//  /c/..///d///")
	fmt.Println(s) // a/b/  /d
}
`



// 获取 path 的绝对路径
Abs(path string) (string, error)
// 判断路径是否为绝对路径
IsAbs(path string) bool
`
func main() {
	s1 := `a/b/c/d`
	fmt.Println(filepath.Abs(s1)) // 不同系统显示不一样
	s2 := `/a/b/c/d`
	fmt.Println(filepath.IsAbs(s1)) // false
	fmt.Println(filepath.IsAbs(s2)) // true
}
`



// 将路径序列 path 分割为多条独立的路径
SplitList(path string) []string
`
func main() {
	path := `a/b/c:d/e/f:   g/h/i`
	s := filepath.SplitList(path)
	fmt.Printf("%q", s)  // ["a/b/c" "d/e/f" "   g/h/i"]
}
`



// 返回路径字符串中的卷名
// Windows 中的 `C:\Windows` 会返回 "C:"
// Linux 中的 `//host/share/name` 会返回 `//host/share`
VolumeName(path string) string
// 返回链接（快捷方式）所指向的实际文件
EvalSymlinks(path string) (string, error)




// 判断 name 是否和指定的模式 pattern 完全匹配
Match(pattern, name string) (matched bool, err error)
// pattern 规则如下：
// 可以使用 ? 匹配单个任意字符（不匹配路径分隔符）。
// 可以使用 * 匹配 0 个或多个任意字符（不匹配路径分隔符）。
// 可以使用 [] 匹配范围内的任意一个字符（可以包含路径分隔符）。
// 可以使用 [^] 匹配范围外的任意一个字符（无需包含路径分隔符）。
// [] 之内可以使用 - 表示一个区间，比如 [a-z] 表示 a-z 之间的任意一个字符。
// 反斜线用来匹配实际的字符，比如 \* 匹配 *，\[ 匹配 [，\a 匹配 a 等等。
// [] 之内可以直接使用 [ * ?，但不能直接使用 ] -，需要用 \]、\- 进行转义。
`
func main() {
	fmt.Println(filepath.Match(`???`,          `abc`))     // true
	fmt.Println(filepath.Match(`???`,          `abcd`))    // false
	fmt.Println(filepath.Match(`*`,            `abc`))     // true
	fmt.Println(filepath.Match(`*`,            ``))        // true
	fmt.Println(filepath.Match(`a*`,           `abc`))     // true
	fmt.Println(filepath.Match(`???\\???`,     `abc\def`)) // true
	fmt.Println(filepath.Match(`???/???`,      `abc/def`)) // true
	fmt.Println(filepath.Match(`/※/※/※/`,      `/a/b/c/`)) // true，　这里用　※　代替　*
	fmt.Println(filepath.Match(`[aA][bB][cC]`, `aBc`))     // true
	fmt.Println(filepath.Match(`[^aA]*`,       `abc`))     // false
	fmt.Println(filepath.Match(`[a-z]*`,       `a+b`))     // true
	fmt.Println(filepath.Match(`\[*\]`,        `[a+b]`))   // true
	fmt.Println(filepath.Match(`[[\]]*[[\]]`,  `[]`))      // true
}
`



// 列出与指定的模式 pattern 完全匹配的文件或目录（匹配原则同上）
Glob(pattern string) (matches []string, err error)
`
func main() {
	// 列出 usr 的子目录中所包含的以 ba（忽略大小写）开头的项目
	list, err := filepath.Glob("/usr/※/[Bb][Aa]*")	//　这里用　※　代替　*
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range list {
		fmt.Println(v)
	}
}
`



// 遍历指定目录（包括子目录），对遍历到的项目用 walkFn 函数进行处理。
Walk(root string, walkFn WalkFunc) error
// 文件处理函数定义如下，如果 WalkFunc 返回 nil，则 Walk 函数继续
// 遍历，如果返回 SkipDir，则 Walk 函数会跳过当前目录（如果当前遍
// 历到的是文件，则同时跳过后续文件及子目录），继续遍历下一个目录。
// 如果返回其它错误，则 Walk 函数会中止遍历过程。
// 在 Walk 遍历过程中，如果遇到错误，则会将错误通过 err 传递给
// WalkFunc 函数，同时 Walk 会跳过出错的项目，继续处理后续项目。
type WalkFunc func(path string, info os.FileInfo, err error) error
`
// WalkFunc 函数：
// 列出含有 *.txt 文件的目录（不是全部，因为会跳过一些子目录）
func findTxtDir(path string, info os.FileInfo, err error) error {
	ok, err := filepath.Match(`*.txt`, info.Name())
	if ok {
		fmt.Println(filepath.Dir(path), info.Name())
		// 遇到 txt 文件则继续处理所在目录的下一个目录
		// 注意会跳过子目录
		return filepath.SkipDir
	}
	return err
}

// WalkFunc 函数：
// 列出所有以 ab 开头的目录（全部，因为没有跳过任何项目）
func findabDir(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		ok, err := filepath.Match(`[aA][bB]*`, info.Name())
		if err != nil {
			return err
		}
		if ok {
			fmt.Println(path)
		}
	}
	return nil
}

func main() {
	// 列出含有 *.txt 文件的目录（不是全部，因为会跳过一些子目录）
	err := filepath.Walk(`/usr`, findTxtDir)
	fmt.Println(err)

	fmt.Println("==============================")

	// 列出所有以 ab 开头的目录（全部，因为没有跳过任何项目）
	err = filepath.Walk(`/usr`, findabDir)
	fmt.Println(err)
}
`
*/

func main() {

}
