package main

import (
	"bytes"
	"fmt"
)

/**
bytes 包 https://www.cnblogs.com/golove/p/3287729.html

对于传入 []byte 的函数，都不会修改传入的参数，返回值要么是参数的副本，要么是参数的切片
*/

// 大小写转换相关
func aboutConversion() {
	fmt.Println()
	fmt.Println()
	fmt.Println("大小写转换相关 ------------------------ start ----------------------")

	var b []byte = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	var B []byte = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}

	// 将切片中的所有内容转成大写
	toUpperRes := bytes.ToUpper(b)
	fmt.Printf("toUpperRes = %s\n", toUpperRes) // toUpperRes = ABCDEFG

	// 将切片中的所有内容转成小写
	toLowerRes := bytes.ToLower(B)
	fmt.Printf("toLowerRes = %s\n", toLowerRes) // toLowerRes = abcdefg

	// 将切片中的所有内容转成标题格式
	toTitleRes := bytes.ToTitle(b)
	fmt.Printf("toTitleRes = %s\n", toTitleRes) // toTitleRes = ABCDEFG

	// 将 s 中的所有单词的首字符修改成为 Title 格式。
	// bug：不能很好的处理已 unicode 标点符号分割的单词。
	titleRes := bytes.Title([]byte("hello world"))
	fmt.Printf("titleRes = %s\n", titleRes) // titleRes = Hello World

	fmt.Println("大小写转换相关 ------------------------ end ----------------------")
}

// 比较相关
// 比较的规则： 引自 https://www.jianshu.com/p/5decb78abe04
// 简单说就是给定两个串，分别从每个串的头开始依次比较串的元素的大小，当地一个不同的元素出现时，
// a[i] != b[i]，比较就结束了，且 a[i] 与 b[i] 的比较结果作为串比较的结果
// 当两个串长度不一样，短串不足的位置会作为空元素处理，且空元素比其他非空元素小。
func aboutCompare() {
	fmt.Println()
	fmt.Println()
	fmt.Println("比较相关 ------------------------ start ----------------------")

	// 比较两个 []byte, nil 参数相当于空的 []byte
	// a < b 返回 -1
	// a == b 返回 0
	// a > b 返回 1
	compareRes := bytes.Compare([]byte{'a'}, []byte{'A', 'B'})
	fmt.Printf("compareRes = %d\n", compareRes) // compareRes = 1

	// 判断两个切片是否相等
	equalRes := bytes.Equal([]byte{'a'}, []byte{'a'})
	fmt.Printf("equalRes = %t\n", equalRes) // equalRes = true

	// 判断两个切片是否相等，忽略大小写和标题三种格式的区别
	equalFoldRes := bytes.EqualFold([]byte{'a'}, []byte{'A'})
	fmt.Printf("equalFoldRes = %t\n", equalFoldRes) // equalFoldRes = true

	fmt.Println("比较相关 ------------------------ end ----------------------")
}

// trim 相关
func aboutTrim() {
	fmt.Println()
	fmt.Println()
	fmt.Println("trim 相关 ------------------------ start ----------------------")

	// 去掉参数一两边包含在参数二中的字符，返回参数一的切片
	trimRes := bytes.Trim([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}, "apology")
	fmt.Printf("trimRes = %s\n", trimRes) // trimRes = bcdef

	// 去掉参数一左边包含在参数二中的字符，返回参数一的切片
	trimLeftRes := bytes.TrimLeft([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}, "apology")
	fmt.Printf("trimLeftRes = %s\n", trimLeftRes) // trimLeftRes = bcdefg

	// 去掉参数一右边包含在参数二中的字符，返回参数一的切片
	trimRightRes := bytes.TrimRight([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}, "apology")
	fmt.Printf("trimRightRes = %s\n", trimRightRes) // trimRightRes = abcdef

	// 去掉参数一两边符合参数二要求的字符，参数二返回 true 的内容会被 trim 掉。函数返回参数一的切片
	trimFuncRes := bytes.TrimFunc([]byte("李白很能喝"), func(r rune) bool {
		return r == '李' || r == '喝'
	})
	fmt.Printf("trimFuncRes = %s\n", trimFuncRes) // trimFuncRes = 白很能

	// 去掉参数一左边符合参数二要求的字符，参数二返回 true 的内容会被 trim 掉。函数返回参数一的切片
	trimLeftFuncRes := bytes.TrimLeftFunc([]byte("杜甫特别忙"), func(r rune) bool {
		return r == '杜'
	})
	fmt.Printf("trimLeftFuncRes = %s\n", trimLeftFuncRes) // trimLeftFuncRes = 甫特别忙

	// 去掉参数一右边符合参数二要求的字符，参数二返回 true 的内容会被 trim 掉。函数返回参数一的切片
	trimRightFuncRes := bytes.TrimRightFunc([]byte("锄禾日当午"), func(r rune) bool {
		return r == '居'
	})
	fmt.Printf("trimRightFuncRes = %s\n", trimRightFuncRes) // trimRightFuncRes = 锄禾日当

	// 去掉左右两边的空格
	trimSpaceRes := bytes.TrimSpace([]byte(" 汗滴禾下土 "))
	fmt.Printf("trimSpaceRes = %q\n", trimSpaceRes) // trimSpaceRes = "汗滴禾下土"

	trimPrefixRes := bytes.TrimPrefix([]byte("谁知盘中餐"), []byte("谁知"))
	fmt.Printf("trimPrefixRes = %s\n", trimPrefixRes) // trimPrefixRes = 盘中餐

	trimSuffixRes := bytes.TrimSuffix([]byte("粒粒皆辛苦"), []byte("辛苦"))
	fmt.Printf("trimSuffixRes = %s\n", trimSuffixRes) // trimSuffixRes = 粒粒皆

	fmt.Println("trim 相关 ------------------------ end ----------------------")
}

// 拆合相关
// 注意：split 开头的函数当 s 的最后正好匹配 sep 的时候，函数返回值里面最后会有一个空串
func aboutSplitAndJoin() {
	fmt.Println()
	fmt.Println()
	fmt.Println("拆合相关 ------------------------ start ----------------------")

	// 以参数二为分隔符将参数一切分成多个子串，结果不包含分隔符。
	// 如果参数二为空，则将参数一切分成 Unicode 字符列表。
	splitRes := bytes.Split([]byte("hello world"), []byte{' ', 'w'})
	for _, v := range splitRes {
		fmt.Printf("splitRes.v = %q\n", v)
	}

	// 作用同 split，但是自己定义切分次数，超过次数的不再切割
	splitNRes := bytes.SplitN([]byte("splitN func test"), []byte{' '}, 2)
	for _, v := range splitNRes {
		fmt.Printf("splitNRes.v = %q\n", v)
	}

	// 功能同 Split，只不过结果包含分隔符（在各个子串尾部）。
	splitAfterRes := bytes.SplitAfter([]byte("split after func test"), []byte{'t'})
	for _, v := range splitAfterRes {
		fmt.Printf("splitAfterRes.v = %q\n", v)
	}

	// 功能同 SplitAfter，只是可以自定义分隔次数
	splitAfterNRes := bytes.SplitAfterN([]byte("split after n func test"), []byte{' '}, 3)
	for _, v := range splitAfterNRes {
		fmt.Printf("splitAfterNRes.v = %q\n", v)
	}

	// Fields 以连续空白为分隔符将 s 切分成多个子串，结果不包含分隔符。
	fieldsRes := bytes.Fields([]byte("hello world"))
	for _, v := range fieldsRes {
		fmt.Printf("fieldsRes.v = %q\n", v)
	}

	// 以符合 f 的字符为分隔符将 s 切分成多个子串，结果不包含分隔符
	fieldsFuncRes := bytes.FieldsFunc([]byte("hello world"), func(r rune) bool {
		return r == 'o'
	})
	for _, v := range fieldsFuncRes {
		fmt.Printf("fieldsFuncRes.v = %q\n", v)
	}

	// 以参数二为连接符将参数一连接成一个字节串
	joinRes := bytes.Join([][]byte{[]byte("hello"), []byte("world")}, []byte{' '})
	fmt.Printf("joinRes = %q\n", joinRes)

	// 将参数一重复参数二次并返回
	repeatRes := bytes.Repeat([]byte("zrh"), 2)
	fmt.Printf("repeatRes = %q\n", repeatRes)

	fmt.Println("拆合相关 ------------------------ end ----------------------")
}

// 子串相关
func aboutSubBytes() {
	fmt.Println()
	fmt.Println()
	fmt.Println("子串相关 ------------------------ start ----------------------")

	// 判断参数一是否以参数二开头
	prefixRes := bytes.HasPrefix([]byte("hello world"), []byte{'h', 'e'})
	fmt.Printf("prefixRes = %t\n", prefixRes)

	// 判断参数一是否以参数二结尾
	suffixRes := bytes.HasSuffix([]byte("hello world"), []byte{'l', 'd'})
	fmt.Printf("suffixRes = %t\n", suffixRes)

	// 判断参数一种是否包含参数二
	containsRes := bytes.Contains([]byte("hello world"), []byte{'l'})
	fmt.Printf("containsRes = %t\n", containsRes)

	// 功能同 Contains，支持中文
	containsRuneRes := bytes.ContainsRune([]byte("明月几时有？把酒问青天。"), '酒')
	fmt.Printf("containsRuneRes = %t\n", containsRuneRes)

	// 参数一中包含参数二中任意一个字符就返回 true，否则返回 false
	containsAnyRes := bytes.ContainsAny([]byte("hello world"), "abcdefg")
	fmt.Printf("containsAnyRes = %t\n", containsAnyRes)

	// 查找参数二在参数一种首次出现的位置，找不到则返回 -1
	indexRes := bytes.Index([]byte("hello world"), []byte{'e', 'l'})
	fmt.Printf("indexRes = %d\n", indexRes) // indexRes = 1

	// 功能同 Index，参数二是一个 byte
	indexByteRes := bytes.IndexByte([]byte("hello world"), 'e')
	fmt.Printf("indexByteRes = %d\n", indexByteRes) // indexByteRes = 1

	// 功能同 Index，参数二是一个 rune
	indexRuneRes := bytes.IndexRune([]byte("不知天上宫阙，今夕是何年。"), '今')
	fmt.Printf("indexRuneRes = %d\n", indexRuneRes) // indexRuneRes = 21，一个中文 3 个字符

	// 查找 chars 中的任何一个字符在 s 中第一次出现的位置，找不到则返回 -1。
	indexAnyRes := bytes.IndexAny([]byte("hello world"), "abcdefg")
	fmt.Printf("indexAnyRes = %d\n", indexAnyRes)

	// 自定义判断方法
	indexFuncRes := bytes.IndexFunc([]byte("hello world"), func(r rune) bool {
		if r == 'a' || r == 'e' || r == 'o' {
			return true
		}
		return false
	})
	fmt.Printf("indexFuncRes = %d\n", indexFuncRes)

	// 返回参数二在参数一中最后一次出现的位置，找不到返回 -1
	lastIndexRes := bytes.LastIndex([]byte("hello world"), []byte{'l'})
	fmt.Printf("lastIndexRes = %d\n", lastIndexRes)

	// 功能同 LastIndex，参数二为 byte
	lastIndexByteRes := bytes.LastIndexByte([]byte("hello world"), 'w')
	fmt.Printf("lastIndexByteRes = %d\n", lastIndexByteRes)

	// 返回参数二中任意一个字符最后一次出现在参数一中的位置，找不到返回 -1
	lastIndexAnyRes := bytes.LastIndexAny([]byte("hello world"), "abcdefg")
	fmt.Printf("lastIndexAnyRes = %d\n", lastIndexAnyRes)

	// 自定义方法查找最后一次出现在参数一中的字符
	lastIndexFuncRes := bytes.LastIndexFunc([]byte("我欲乘风归去，又恐琼楼玉宇，高处不胜寒。起舞弄清影，何似在人间。"), func(r rune) bool {
		if r == '去' || r == '人' {
			return true
		}
		return false
	})
	fmt.Printf("lastIndexFuncRes = %d\n", lastIndexFuncRes)

	// 获取 sep 在 s 中出现的次数（sep 不能重叠）。
	countRes := bytes.Count([]byte("hello world"), []byte{'l'})
	fmt.Printf("countRes = %d\n", countRes)

	fmt.Println("子串相关 ------------------------ end ----------------------")
}

// 替换相关
func aboutReplace() {
	fmt.Println()
	fmt.Println()
	fmt.Println("替换相关 ------------------------ start ----------------------")

	// 将 s 中前 n 个 old 替换为 new，n < 0 则替换全部。
	replaceRes := bytes.Replace([]byte("hello world"), []byte{'l'}, []byte{'z'}, 2)
	fmt.Printf("replaceRes = %q\n", replaceRes)

	// 匹配到的全部替换
	replaceAllRes := bytes.ReplaceAll([]byte("hello world"), []byte{'l'}, []byte{'z'})
	fmt.Printf("replaceAllRes = %q\n", replaceAllRes)

	// 将 s 中的字符替换为 mapping(r) 的返回值，
	// 如果 mapping 返回负值，则丢弃该字符。
	mapRes := bytes.Map(func(r rune) rune {
		if r == 'e' {
			return 'f'
		}
		if r == 'd' {
			return 'p'
		}
		if r == 'o' {
			return -1
		}
		return r
	}, []byte("hello world"))
	fmt.Printf("mapRes = %q\n", mapRes) // mapRes = "hfll wrlp"

	// 将 s 转换为 []rune 类型返回
	runesRes := bytes.Runes([]byte("hello 世界"))
	fmt.Printf("runesRes = %q\n", runesRes)

	fmt.Println("替换相关 ------------------------ end ----------------------")
}

func aboutBytes() {
	// 转换相关
	aboutConversion()

	// 比较相关
	aboutCompare()

	// trim 相关
	aboutTrim()

	// 拆合相关
	aboutSplitAndJoin()

	// 子串相关
	aboutSubBytes()

	// 替换相关
	aboutReplace()
}
