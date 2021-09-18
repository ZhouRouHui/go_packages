package main

import (
	"fmt"
	"reflect"
)

// https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/%E5%8F%8D%E5%B0%84.html
/**
1. 反射
反射是指在程序运行期间对程序本身进行访问和修改的能力

1.1.1 变量的内在机制
* 变量包含类型信息和值信息  var arr [10]int arr[0] = 10
* 类型信息：是静态的元信息，是预先定义好的
* 值信息：使程序运行过程中动态改变的

1.1.2 反射的作用
* reflect 包封装了反射相关的方法
* 获取类型信息: reflect.TypeOf, 是静态的
* 获取值信息: reflect.ValueOf, 是动态的

1.1.3 空接口与反射
* 反射可以在运行时动态获取程序的各种详细信息
* 反射获取 interface 类型信息
`
package main

import (
   "fmt"
   "reflect"
)

//反射获取interface类型信息

func reflect_type(a interface{}) {
   t := reflect.TypeOf(a)
   fmt.Println("类型是：", t)
   // kind()可以获取具体类型
   k := t.Kind()
   fmt.Println(k)
   switch k {
   case reflect.Float64:
      fmt.Printf("a is float64\n")
   case reflect.String:
      fmt.Println("string")
   }
}

func main() {
   var x float64 = 3.4
   reflect_type(x)
}
`

* 反射获取 interface 值信息
`
package main

import (
    "fmt"
    "reflect"
)

//反射获取interface值信息

func reflect_value(a interface{}) {
    v := reflect.ValueOf(a)
    fmt.Println(v)
    k := v.Kind()
    fmt.Println(k)
    switch k {
    case reflect.Float64:
        fmt.Println("a是：", v.Float())
    }
}

func main() {
    var x float64 = 3.4
    reflect_value(x)
}
`

* 反射修改值信息
`
package main

import (
    "fmt"
    "reflect"
)

//反射修改值
func reflect_set_value(a interface{}) {
    v := reflect.ValueOf(a)
    k := v.Kind()
    switch k {
    case reflect.Float64:
        // 反射修改值
        v.SetFloat(6.9)
        fmt.Println("a is ", v.Float())
    case reflect.Ptr:
        // Elem()获取地址指向的值
        v.Elem().SetFloat(7.9)
        fmt.Println("case:", v.Elem().Float())
        // 地址
        fmt.Println(v.Pointer())
    }
}

func main() {
    var x float64 = 3.4
    // 反射认为下面是指针类型，不是float类型
    reflect_set_value(&x)
    fmt.Println("main:", x)
}
`

1.1.4 结构体与反射
查看类型、字段、方法
`
package main

import (
    "fmt"
    "reflect"
)

// 定义结构体
type User struct {
    Id   int
    Name string
    Age  int
}

// 绑方法
func (u User) Hello() {
    fmt.Println("Hello")
}

// 传入interface{}
func Poni(o interface{}) {
    t := reflect.TypeOf(o)
    fmt.Println("类型：", t)
    fmt.Println("字符串类型：", t.Name())
    // 获取值
    v := reflect.ValueOf(o)
    fmt.Println(v)
    // 可以获取所有属性
    // 获取结构体字段个数：t.NumField()
    for i := 0; i < t.NumField(); i++ {
        // 取每个字段
        f := t.Field(i)
        fmt.Printf("%s : %v", f.Name, f.Type)
        // 获取字段的值信息
        // Interface()：获取字段对应的值
        val := v.Field(i).Interface()
        fmt.Println("val :", val)
    }
    fmt.Println("=================方法====================")
    for i := 0; i < t.NumMethod(); i++ {
        m := t.Method(i)
        fmt.Println(m.Name)
        fmt.Println(m.Type)
    }

}

func main() {
    u := User{1, "zs", 20}
    Poni(u)
}
`

* 查看匿名字段
`
package main

import (
    "fmt"
    "reflect"
)

// 定义结构体
type User struct {
    Id   int
    Name string
    Age  int
}

// 匿名字段
type Boy struct {
    User
    Addr string
}

func main() {
    m := Boy{User{1, "zs", 20}, "bj"}
    t := reflect.TypeOf(m)
    fmt.Println(t)
    // Anonymous：匿名
    fmt.Printf("%#v\n", t.Field(0))
    // 值信息
    fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}
`

* 修改结构体的值
`
package main

import (
    "fmt"
    "reflect"
)

// 定义结构体
type User struct {
    Id   int
    Name string
    Age  int
}

// 修改结构体值
func SetValue(o interface{}) {
    v := reflect.ValueOf(o)
    // 获取指针指向的元素
    v = v.Elem()
    // 取字段
    f := v.FieldByName("Name")
    if f.Kind() == reflect.String {
        f.SetString("kuteng")
    }
}

func main() {
    u := User{1, "5lmh.com", 20}
    SetValue(&u)
    fmt.Println(u)
}
`

* 调用方法
`
package main

import (
    "fmt"
    "reflect"
)

// 定义结构体
type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) Hello(name string) {
    fmt.Println("Hello：", name)
}

func main() {
    u := User{1, "5lmh.com", 20}
    v := reflect.ValueOf(u)
    // 获取方法
    m := v.MethodByName("Hello")
    // 构建一些参数
    args := []reflect.Value{reflect.ValueOf("6666")}
    // 没参数的情况下：var args2 []reflect.Value
    // 调用方法，需要传入方法的参数
    m.Call(args)
}
`

* 获取字段的tag
`
package main

import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string `json:"name1" db:"name2"`
}

func main() {
    var s Student
    v := reflect.ValueOf(&s)
    // 类型
    t := v.Type()
    // 获取字段
    f := t.Elem().Field(0)
    fmt.Println(f.Tag.Get("json"))
    fmt.Println(f.Tag.Get("db"))
}
`
*/

// ReflectType 反射获取 interface 类型信息
func ReflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是: ", t)

	// Kind() 可以获取具体类型
	k := t.Kind()
	fmt.Printf("具体类型是 %s\n", k)

	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}
}

// 测试反射获取 interface 类型信息
func testReflectType() {
	var x float64 = 3.4
	ReflectType(x)
}

// ReflectValue 反射获取 interface 值信息
func ReflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Printf("a 的值是 %s\n", v.String())

	// 获取值的类型
	k := v.Kind()
	fmt.Printf("a 的类型是 %s\n", k)

	switch k {
	case reflect.Float64:
		// v.Float() 以 float64 形式返回 v 的基础值。
		fmt.Println("a 是 ", v.Float())
	}
}

// 反射获取 interface 值信息
func testReflectValue() {
	var x float64 = 3.4
	ReflectValue(x)
}

// ReflectSetValue 反射修改值
func ReflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem() 获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println("地址:", v.Pointer())
	}
}

// 反射修改值
func testReflectSetValue() {
	var x float64 = 3.4
	// 反射认为下面是指针类型，不是 float 类型
	ReflectSetValue(&x)
	fmt.Println("main:", x)
}

// User 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// Hello 绑方法
func (u User) Hello(name string) {
	fmt.Println("Hello", name)
}

// Poni 传入 interface{}
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型:", t)
	fmt.Println("字符串类型:", t.Name())

	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println("值:", v)

	// 可以获取所有属性，获取结构体的字段个数 t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 获取每个字段
		f := t.Field(i)
		fmt.Printf("字段名：%s, 字段类型：%v\n", f.Name, f.Type)

		// 获取字段的值信息 Interface() ：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Printf("字段 %s 的值为 %v\n", f.Name, val)
	}

	fmt.Println("================= 方法 ==================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("方法的名称: ", m.Name)
		fmt.Println("方法的类型", m.Type)
	}
}

// 查看结构体的类型、字段和方法
func testPoni() {
	u := User{1, "zs", 20}
	Poni(u)
}

// Boy 匿名字段
type Boy struct {
	User
	Addr string
}

// 测试查看匿名字段
func testAnonymousField() {
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

// SetValue 修改结构体值
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
}

// testSetValue 测试修改结构体字段的值
func testSetValue() {
	u := User{1, "51mh.com", 20}
	// 修改值要传指针
	SetValue(&u)
	fmt.Println(u)
}

// 测试调用结构体方法
func testCallMethod() {
	u := User{1, "5lmh.com", 20}
	v := reflect.ValueOf(u)
	// 获取方法
	m := v.MethodByName("Hello")
	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf("6666")}
	// 没参数的情况下：var args2 []reflect.Value
	// 调用方法，需要传入方法的参数
	m.Call(args)
}

type Student struct {
	Name string `json:"name1" db:"name2"`
}

// 测试获取结构体字段的 tag
func testGetTag() {
	var s Student
	v := reflect.ValueOf(&s)
	// 类型
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}

func main() {
	// 测试反射获取 interface 类型信息
	testReflectType()
	fmt.Println()

	// 反射获取 interface 值信息
	testReflectValue()
	fmt.Println()

	// 反射修改值
	testReflectSetValue()
	fmt.Println()

	// 查看结构体的类型、字段和方法
	testPoni()
	fmt.Println()

	// 测试查看匿名字段
	testAnonymousField()
	fmt.Println()

	// 测试修改结构体字段的值
	testSetValue()
	fmt.Println()

	// 测试调用结构体方法
	testCallMethod()
	fmt.Println()

	// 测试获取结构体字段的 tag
	testGetTag()
	fmt.Println()
}
