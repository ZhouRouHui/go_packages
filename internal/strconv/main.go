package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	字符串数字转 int
	 */
	fmt.Println("字符串数字转 int ------------------------- start -----------------------------")
	// 将字符串转成数字
	// 参数一：数字字符串
	// 参数二：数字字符串的进制，二进制，八进制，十进制，十六进制
	// 参数三：返回结果的 bit 大小，也就是 int8 int16 int32 int64
	parseIntRes, _ := strconv.ParseInt("-123", 10, 64)
	fmt.Printf("parseIntRes = %d\n", parseIntRes)

	// 作用与 ParseInt 相同，但是 ParseUint 是无符号的，如果参数一是小于 0 的内容，将返回 0
	parseUintRes, _ := strconv.ParseUint("-124", 10, 32)
	fmt.Printf("parseUintRes = %d\n", parseUintRes)

	// 字符串数字转换成 int，等同于 ParseInt(s, 10, 0), 返回 int 类型
	atoiRes, _ := strconv.Atoi("-12")
	fmt.Printf("atoiRes = %d\n", atoiRes)

	fmt.Println("字符串数字转 int ------------------------- end -----------------------------")



	/**
	int 转字符串
	*/
	fmt.Println()
	fmt.Println()
	fmt.Println("int 转字符串 ------------------------- start -----------------------------")

	formatIntRes := strconv.FormatInt(-34, 10)
	fmt.Printf("formatIntRes = %s\n", formatIntRes)

	formatUintRes := strconv.FormatUint(34, 10)
	fmt.Printf("formatUintRes = %s\n", formatUintRes)

	itoaRes := strconv.Itoa(99)
	fmt.Printf("itoaRes = %s\n", itoaRes)

	// 将参数二转成字符串并追加到参数一中，参数三为参数二的进制
	appendIntRes := strconv.AppendInt([]byte("append int func test "), -123, 10)
	fmt.Printf("appendIntRes = %s\n", string(appendIntRes))

	// 作用同 AppendInt, 但是无符号
	appendUintRes := strconv.AppendUint([]byte("append uint func test "), 123, 10)
	fmt.Printf("appendUintRes = %s\n", appendUintRes)

	fmt.Println("int 转字符串 ------------------------- end -----------------------------")



	/**
	布尔型操作
	 */
	fmt.Println()
	fmt.Println()
	fmt.Println("布尔型操作 ------------------------- start -----------------------------")

	// 参数支持 "1", "t", "T", "true", "TRUE", "True", "0", "f", "F", "false", "FALSE", "False"
	// 其他字符串传进去会返回 err
	parseBoolRes, err := strconv.ParseBool("f")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parseBoolRes = %t\n", parseBoolRes)

	// 返回参数的字符串形式 "true" 或 "false"
	formatBoolRes := strconv.FormatBool(true)
	fmt.Printf("formatBoolRes = %s\n", formatBoolRes)

	// 根据参数二的值转成字符串形式，并添加到参数一里面
	appendBoolRes := strconv.AppendBool([]byte("hello world "), true)
	fmt.Printf("appendBoolRes = %s\n", string(appendBoolRes))

	fmt.Println("布尔型操作 ------------------------- end -----------------------------")



	/**
	浮点型操作
	*/
	fmt.Println()
	fmt.Println()
	fmt.Println("浮点型操作 ------------------------- start -----------------------------")

	// 字符串转浮点型，参数二为 bit 位数
	parseFloatRes, _ := strconv.ParseFloat("1.2", 64)
	fmt.Printf("parseFloatRes = %f\n", parseFloatRes)


	// FormatFloat 将浮点数 f 转换为字符串值
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64）
	//
	// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	//
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	formatFloatRes := strconv.FormatFloat(2.3, 'e', 2, 64)
	fmt.Printf("formatFloatRes = %s\n", formatFloatRes)


	// 将浮点数转成 string 类型，并追加到参数一中，后面三个参数和 FormatFloat 规则一样
	appendFloatRes := strconv.AppendFloat([]byte("append float func test "), 2.3, 'e', 2, 64)
	fmt.Printf("appendFloatRes = %s\n", string(appendFloatRes))

	fmt.Println("浮点型操作 ------------------------- end -----------------------------")


	fmt.Println()

	// 给字符串添加前后双引号
	quoteRes := strconv.Quote("quote test")
	fmt.Printf("quoteRes = %s\n", quoteRes)
}
