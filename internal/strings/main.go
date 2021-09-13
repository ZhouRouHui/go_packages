package main

import (
	"fmt"
	"strings"
)

func main() {

	/**
	index 相关
	 */
	fmt.Println()
	fmt.Println()
	fmt.Println("index 相关函数 --------------------- start --------------------")

	// 返回参数一中参数二第一次出现的位置下标，如果不存在则返回 -1
	indexRes := strings.Index("index func test", "l")
	fmt.Printf("indexRes = %d\n", indexRes)

	// 返回参数一中参数二第一次出现的位置下标，参数二为 byte 类型，如果不存在则返回 -1
	indexByteRes := strings.IndexByte("index byte func test", 'b')
	fmt.Printf("indexByteRes = %d\n", indexByteRes)

	// 返回参数一中参数二第一次出现的位置下标，参数二支持中文字符，如果不存在则返回 -1
	indexRuneRes := strings.IndexRune("index rune func test 你", '你')
	fmt.Printf("indexRuneRes = %d\n", indexRuneRes)

	// 返回参数一中参数二任意一个字符第一次出现的位置下标，如果都不存在则返回 -1
	indexAnyRes := strings.IndexAny("index any func test", "abci")
	fmt.Printf("indexAnyRes = %d\n", indexAnyRes)

	// 返回参数一中参数二最后一次出现的位置下标，如果不存在则返回 -1
	lastIndexRes := strings.LastIndex("last index func test", "st")
	fmt.Printf("lastIndexRes = %d\n", lastIndexRes)

	// 返回参数一中参数二任意一个字符最后出现的一个位置，以下标最大的为值，如果都不存在则返回 -1
	lastIndexAnyRes := strings.LastIndexAny("last index any func test", "st")
	fmt.Printf("lastIndexAnyRes = %d\n", lastIndexAnyRes)

	// 返回参数一中参数二最后一次出现的位置下标，参数二为 byte 类型，如果不存在则返回 -1
	lastIndexByteRes := strings.LastIndexByte("last index byte test", 'a')
	fmt.Printf("lastIndexByteRes = %d\n", lastIndexByteRes)

	fmt.Println("index 相关函数 --------------------- end --------------------")


	/**
	contains 相关
	 */
	fmt.Println()
	fmt.Println()
	fmt.Println("contains 相关函数 --------------------- start --------------------")

	// 判断参数一中是否包含第二个参数，返回值布尔型
	containsRes := strings.Contains("contains func test", "con")
	fmt.Printf("containsRes = %t\n", containsRes)

	// 判断字符串参数一中是否包含第二个参数中的任意一个字符，包含则返回 true
	// 如果第二个参数为空字符串则返回 false
	containsAnyRes := strings.ContainsAny("contains any func test", "")
	fmt.Printf("containsAnyRes = %t\n", containsAnyRes)

	// 判断参数一中是否包含参数二，参数二为字符并且支持中文
	containsRuneRes := strings.ContainsRune("contains rune func test 你好", '中')
	fmt.Printf("containsRuneRes = %t\n", containsRuneRes)

	fmt.Println("contains 相关函数 --------------------- end --------------------")


	/**
	split 相关函数
	 */
	fmt.Println()
	fmt.Println()
	fmt.Println("split 相关函数 --------------------- start --------------------")

	// 返回参数一按参数二分割后的切片
	splitRes := strings.Split("split func test", " ")
	fmt.Printf("splitRes = %v, len of splitRes = %d\n", splitRes, len(splitRes))
	for _, v := range splitRes {
		fmt.Printf("splitRes.v = %s\n", v)
	}

	fmt.Println()

	// 返回将参数一按参数二分割后的切片，参数三控制返回值得内容
	// 参数三 > 0 时，切片长度最大为参数三的值
	// 参数三 == 0 时，切片为 nil（零子串 ）
	// 参数三 < 0 时，切片为参数一按照参数二分割后的所有子串的长度
	splitNRes := strings.SplitN("split n func test", " ", 3)
	fmt.Printf("splitNRes = %v, len of splitNRes = %d\n", splitNRes, len(splitNRes))
	for _, v := range splitNRes {
		fmt.Printf("splitNRes.v = %s\n", v)
	}

	fmt.Println()

	// 返回参数一按参数二分割后的切片，切片元素包含参数二本身
	splitAfterRes := strings.SplitAfter("split zafter zfunc ztest", "z")
	fmt.Printf("splitAfterRes = %v, len of splitAfterRes = %d\n", splitAfterRes, len(splitAfterRes))
	for _, v := range splitAfterRes {
		fmt.Printf("splitAfterRes.v = %s\n", v)
	}

	fmt.Println()

	// 返回将参数一按参数二分割后的切片，切片中的元素包含参数二本身，参数三控制返回值内容，规则同 splitN 的参数三
	splitAfterNRes := strings.SplitAfterN("split zafter zn zfunc ztest", "z", -1)
	fmt.Printf("splitAfterNRes = %v, len of splitAfterNRes = %d\n", splitAfterNRes, len(splitAfterNRes))
	for _, v := range splitAfterNRes {
		fmt.Printf("splitAfterN.v = %s\n", v)
	}

	fmt.Println()

	// 以连续的空白字符为分隔符，将参数一切分成多个子串，结果中不包含空白字符本身。
	// 空白字符有: \t \n \v \f \r U+0085 (NEL), U+00A0 (NBSP)
	// 如果参数一中只包含空白字符，则返回一个空列表
	fieldsRes := strings.Fields("fields func test")
	fmt.Printf("fieldsRes = %v, len of fieldRes = %d\n", fieldsRes, len(fieldsRes))
	for _, v := range fieldsRes {
		fmt.Printf("fieldsRes.v = %v\n", v)
	}

	fmt.Println()

	fmt.Println("split 相关函数 --------------------- end --------------------")


	/**
	trim 相关函数
	*/
	fmt.Println()
	fmt.Println()
	fmt.Println("trim 相关函数 --------------------- start --------------------")
	// 删除参数一中头尾的参数二的内容
	trimRes := strings.Trim("trim func test", "t")
	fmt.Printf("trimRes = %s\n", trimRes)

	// 删除参数一头部的参数二的内容
	trimLeftRes := strings.TrimLeft("trim left func test", "t")
	fmt.Printf("trimLeftRes = %s\n", trimLeftRes)

	// 删除参数一尾部的参数二的内容
	trimRightRes := strings.TrimRight("trim right func test", "t")
	fmt.Printf("trimRightRes = %s\n", trimRightRes)

	// 删除参数收尾连续的空白字符
	trimSpaceRes := strings.TrimSpace("trim space func test")
	fmt.Printf("trimSpaceRes = %s\n", trimSpaceRes)

	// 如果参数一以参数二开头，则将参数一中头部的参数二的内容删除后返回。
	// 如果参数一头部不是以参数二开头，则原样返回
	trimPrefixRes := strings.TrimPrefix("trim prefix func test", "tr")
	fmt.Printf("trimPrefixRes = %s\n", trimPrefixRes)

	// 如果参数一以参数二开头，则将参数一中头部的参数二的内容删除后返回。
	// 如果参数一头部不是以参数二开头，则原样返回
	trimSuffixRes := strings.TrimSuffix("trim suffix func test", "st")
	fmt.Printf("trimSuffixRes = %s\n", trimSuffixRes)

	fmt.Println("trim 相关函数 --------------------- end --------------------")


	/**
	其他函数
	*/
	fmt.Println()
	fmt.Println()
	fmt.Println("其他函数 --------------------- start --------------------")

	// 字符串替换，将参数一中的参数二用参数三替换，参数四为替换次数，如果参数四小于 0，则将匹配到的全部替换
	replaceRes := strings.Replace("replace func test", "e", "z", -1)
	fmt.Printf("replaceRes = %s\n", replaceRes)

	// 将参数一中的参数二用参数三替换
	replaceAllRes := strings.ReplaceAll("replace all func test", "e", "z")
	fmt.Printf("replaceAllRes = %s\n", replaceAllRes)

	hasPrefixRes := strings.HasPrefix("has prefix func test", "has")
	fmt.Printf("hasPrefixRes = %t\n", hasPrefixRes)

	hasSuffixRes := strings.HasSuffix("has suffix func test", "has")
	fmt.Printf("hasSuffixRes = %t\n", hasSuffixRes)

	// 将参数一的元素以参数二为连接符，连接成一个字符串
	joinRes := strings.Join([]string{"join", "func", "test"}, " ")
	fmt.Printf("joinRes = %s\n", joinRes)

	// 返回参数一中参数二出现的次数，如果参数二为空字符串，则返回参数一中Unicode字符数量 + 1
	countRes := strings.Count("count func test", "")
	fmt.Printf("countRes = %d\n", countRes)

	// 字符串忽略大小写的比较
	equalFoldRes := strings.EqualFold("EQUAL fold FUNC test", "equal FOLD func TEST")
	fmt.Printf("equalFoldRes = %t\n", equalFoldRes)

	// 将参数一重复参数二的次数，并返回
	repeatRes := strings.Repeat("repeat func test ", 2)
	fmt.Printf("repeatRes = %s\n", repeatRes)

	// 将字符串转成大写
	toUpperRes := strings.ToUpper("to upper func test")
	fmt.Printf("toUpperRes = %s\n", toUpperRes)

	// 将字符串转成小写
	toLowerRes := strings.ToLower("TO LOWER FUNC TEST")
	fmt.Printf("toLowerRes = %s\n", toLowerRes)

	// 将参数中的所有字符修改为其 Title 格式，大部分字符的 Title 格式就是 Upper 格式，
	// 只有少数字符的 Title 格式是特殊字符。这里的 ToTitle 主要给 Title 函数调用
	toTitleRes := strings.ToTitle("to TITLE func TEST")
	fmt.Printf("toTitleRes = %s\n", toTitleRes)

	// 返回参数二的副本，参数一会将参数二中的每一个字符做处理，当某一个字符在参数一中返回负值，则当前字符被删除。
	mapRes := strings.Map(func(r rune) rune {
		if r == 'a' {
			return -1
		}
		if r == 'b'{
			r = 'z'
		}
		return r
	}, "abc")
	fmt.Printf("mapRes = %s\n", mapRes)

	fmt.Println("其他函数 --------------------- end --------------------")
}
