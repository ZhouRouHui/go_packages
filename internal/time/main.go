package main

import (
	"fmt"
	"runtime"
	"time"
)

// TimeLayout 官方的时间 layout 格式
const TimeLayout = "2006-01-02 15:04:05"

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("time.Now() = %v\n", now)

	// 用具体值生成一个 time.Time 类型的返回值
	dataRes := time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)
	fmt.Println(dataRes)


	/**
	Time 类型相关函数
	 */
	testTime()



	/**
	Parse 相关函数
	 */
	//testParse()


	/**
	Timer 一次性的定时器
	*/
	//timerTest()


	/**
	Ticker 周期性定时器
	循环周期性的往 Ticker.C 里面发送时间，每次的时间为当前时间加上参数
	如：
		2000-01-01 00:00:00
		2000-01-01 00:00:03
		2000-01-01 00:00:06
		2000-01-01 00:00:09
	*/
	//testTicker()
}

// testTime Time 类型相关函数
func testTime() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Time 类型相关函数 -------------------------- start ------------------------------")

	timeObj := time.Now()
	fmt.Printf("timeObj = %v\n", timeObj)

	// 将一个非 UTC 的 Time 转成 UTC 格式
	utcTimeObj := timeObj.UTC()
	fmt.Printf("utcTimeObj = %v\n", utcTimeObj)

	// 将一个非 Local 的 Time 转成 Local
	timeObj = utcTimeObj.Local()
	fmt.Printf("timeObj = %v\n", timeObj)



	/**
	获取时间对象里面的内容
	 */
	// 解析一个时间的对象，返回里面的 年，月，日 三个参数
	year, month, day := timeObj.Date()
	fmt.Printf("year = %d, month = %d, day = %d\n", year, month, day)

	// 返回时间对象的 时，分，秒 三个参数
	hour, min, sec := timeObj.Clock()
	fmt.Printf("hour = %d, min = %d, sec = %d\n", hour, min, sec)

	// 返回时间对象的年的值
	yearRes := timeObj.Year()
	fmt.Printf("yearRes = %d\n", yearRes)

	// 返回时间对象的月的值
	monthRes := timeObj.Month()
	fmt.Printf("monthRes = %d\n", monthRes)

	// 返回时间对象的日期的值
	dayRes := timeObj.Day()
	fmt.Printf("dayRes = %d\n", dayRes)

	// 解析对象，返回小时的值
	hourRes := timeObj.Hour()
	fmt.Printf("hourRes = %d\n", hourRes)

	// 返回时间对象的分钟数的值
	minuteRes := timeObj.Minute()
	fmt.Printf("minuteRes = %d\n", minuteRes)

	// 返回一个 Time 的秒数
	secondRes := timeObj.Second()
	fmt.Printf("secondRes = %v\n", secondRes)

	// 返回一个 Time 的纳秒数
	nanosecondRes := timeObj.Nanosecond()
	fmt.Printf("nanosecondRes = %d\n", nanosecondRes)

	// 返回对象的时区
	locationRes := timeObj.Location()
	fmt.Printf("locationRes = %v\n", locationRes)

	// 返回对象的时间戳格式
	unixRes := timeObj.Unix()
	fmt.Printf("unixRes = %d\n", unixRes)

	// 返回对象的时间戳纳秒格式
	unixNanoRes := timeObj.UnixNano()
	fmt.Printf("unixNanoRes = %d\n", unixNanoRes)

	// 返回对象当前的星期几，注：星期日的值为 0
	weekdayRes := timeObj.Weekday()
	fmt.Printf("weekdayRes = %d\n", weekdayRes)

	// 返回时间对象中当前日期在当年中的第几天
	yearDayRes := timeObj.YearDay()
	fmt.Printf("yearDayRes = %d\n", yearDayRes)



	/**
	两个时间对象比较
	 */
	// 判断两个时间对象是否相等
	equalRes := timeObj.Equal(time.Date(2021, 1, 1, 1, 1, 1, 1, time.Local))
	fmt.Printf("equalRes = %t\n", equalRes)

	// 判断当前时间对象是否在另一个时间对象之前
	beforeRes := timeObj.Before(time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local))
	fmt.Printf("beforeRes = %v\n", beforeRes)

	// 判断一个时间对象是否在另一个时间对象之后
	afterRes := timeObj.After(time.Date(2022, 1, 1, 1, 1, 1, 1, time.Local))
	fmt.Printf("afterRes = %t\n", afterRes)



	/**
	格式化
	 */
	// 将时间对象按给定的格式返回字符串
	formatRes := timeObj.Format("2006/01/02")
	fmt.Printf("formatRes = %s\n", formatRes)

	// 将一个时间对象转成 string
	stringRes := timeObj.String()
	fmt.Printf("stringRes = %s\n", stringRes)



	/**
	计算
	 */
	// 将时间对象添加一个时间段生成新的时间对象，参数支持负数，为负数时表示减掉一个时间段
	addRes := timeObj.Add(-1 * time.Hour)
	fmt.Printf("addRes = %v\n", addRes)

	// 按年，月，日添加时间，支持负数，返回一个计算之后的时间对象
	addDateRes := timeObj.AddDate(-2, 2, 2)
	fmt.Printf("addDateRes = %v\n", addDateRes)

	// 当前时间减去另一个时间对象，返回一个 time.Duration
	subRes := timeObj.Sub(time.Date(2021, 9, 10, 0, 0, 0, 0, time.Local))
	fmt.Printf("subRes = %v\n", subRes)




	// 如果当前对象是一个零时间，返回 true 否则返回 false
	zeroObj := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	isZeroRes := zeroObj.IsZero()
	fmt.Printf("isZeroRes = %t\n", isZeroRes)


	fmt.Println("Time 类型相关函数 -------------------------- end ------------------------------")
}

// testParse parse 相关函数
func testParse() {
	fmt.Println("Parse 相关函数 ------------------------ start -----------------------------")

	// 将一个时间的字符串转换成 time.Time 类型
	parseRes, _ := time.Parse(TimeLayout, "2021-09-19 13:18:13")
	fmt.Printf("parseRes = %v\n", parseRes)

	// 功能与 Parse 相似，参数三为时区
	parseInLocationRes, _ := time.ParseInLocation(TimeLayout, "2021-09-19 13:18:13", time.Local)
	fmt.Printf("parseInLocationRes = %v\n", parseInLocationRes)

	// 将参数解析成 time.Duration 类型，参数支持 "ns", "us" (or "µs"), "ms", "s", "m", "h" 这些单位
	parseDurationRes, _ := time.ParseDuration("10h")
	fmt.Printf("parseDurationRes = %v\n", parseDurationRes)

	// 将当前 goroutine 暂停一段时间
	time.Sleep(time.Second)

	fmt.Println("Parse 相关函数 ------------------------ end -----------------------------")
}

// timerTest 一次性定时器测试
func timerTest() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Timer 定时器 -------------------------- start ------------------------------")

	// 测试创建一个定时器
	testCreateTimer()
	fmt.Println()

	// 测试重置定时器的时间
	testResetTimer()
	fmt.Println()

	// 测试将定时器停止
	testStopTimer()
	fmt.Println()

	// After 返回一个 Time 类型的 channel，值为当前时间加上参数给的时间
	fmt.Println(time.Now())
	afterRes := time.After(10 * time.Second)
	fmt.Println(<-afterRes)

	fmt.Println("Timer 一次性的定时器 -------------------------- end ------------------------------")
}

// testCreateTimer 测试创建一个定时器
func testCreateTimer() {
	// 定时器生成时，里面的 C 的时间就是当前时间加上参数的值，就算中间有 sleep 也不会影响
	newTimerRes := time.NewTimer(time.Second * 3)
	fmt.Println("当前时间为：", time.Now())
	time.Sleep(5 * time.Second)	// 故意 sleep 10秒，查看结果
	t := <-newTimerRes.C	// 从定时器拿数据
	fmt.Println("当前时间为：", t)
}

// testResetTimer 测试重置定时器的时间
func testResetTimer() {
	// 使用 Reset 重置定时器的时间
	timerReset := time.NewTimer(time.Second * 10)
	go timerGoroutine(timerReset)
	// 在这里把定时器的时间充值为 3秒
	timerReset.Reset(3 * time.Second)
	time.Sleep(5 * time.Second)	// 这里要 sleep 一下，否则新起的 goroutine 无法执行完成
}

// testStopTimer 测试将定时器停止
func testStopTimer() {
	// 使用 Stop 将定时器停止
	timerStop := time.NewTimer(time.Second * 3)
	go timerGoroutine(timerStop)
	timerStop.Stop()	// stop 后,上面的 goroutine 中将无法从 chan 里面取得时间
	time.Sleep(5 * time.Second)
}

// timerGoroutine 定时器的测试 goroutine
func timerGoroutine(timer *time.Timer) {
	fmt.Println("timer goroutine --- start ---")
	fmt.Println("当前时间：", time.Now())
	t := <-timer.C
	fmt.Println("当前时间：", t)
	fmt.Println("timer goroutine --- end ---")
}

// testTicker 周期性定时器测试
func testTicker() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Ticker 周期性定时器 -------------------------- start ------------------------------")

	// 测试创建一个周期性的定时器
	testCreateTicker()


	// 测试停止周期性的定时器
	testTickerStop()

	fmt.Println("Ticker 周期性定时器 -------------------------- end ------------------------------")
}

// testCreateTicker 测试创建一个周期性的定时器
func testCreateTicker() {
	// NewTicker 创建一个周期性的定时器
	tickerObj := time.NewTicker(3 * time.Second)
	fmt.Println("当前时间为：", time.Now())
	go func() {
		for {
			fmt.Println("ticker goroutine for -----------------")
			t := <-tickerObj.C	// 从定时器中获取数据
			fmt.Println("当前时间为：", t)
			fmt.Println()
		}
	}()
	time.Sleep(15 * time.Second)
}

// testTickerStop 测试停止周期性的定时器
func testTickerStop() {
	// Stop 停止周期性的定时器
	tickerStopObj := time.NewTicker(3 * time.Second)
	count := 1
	fmt.Println("当前时间：", time.Now(), "计数器 = ", count)
	go func() {
		for {
			fmt.Println("ticker stop for -----------------")
			t := <- tickerStopObj.C
			count++
			fmt.Println("当前时间：", t, "计数器 = ", count)
			fmt.Println()

			if count == 5 {
				// 将定时器停止
				tickerStopObj.Stop()
				// 退出当前 goroutine，交出控制权
				runtime.Goexit()
			}
		}
	}()
	time.Sleep(15 * time.Second)
}
