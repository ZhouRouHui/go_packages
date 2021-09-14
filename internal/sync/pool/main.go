package main

import (
	"fmt"
	"sync"
)

/**
Pool 池子

这个池子的目的就是为了复用已经用过的对象，来达到优化内存使用和回收的目的。
当对象特别大并且使用频繁的时候，大大的减少对象的创建和回收的时间。

* 通过 New 去定义你这个池子里面放的究竟是什么东西，在这个池子里面你只能放一种类型的东西。比如在上面的例子中我就在池子里面放了字符串。
* 我们随时可以通过 Get 方法从池子里面获取我们之前在 New 里面定义类型的数据。
* 当我们用完了之后可以通过 Put 方法放回去，或者放别的同类型的数据进去。
*/

var p sync.Pool = sync.Pool{
	New: func() interface{} {
		return "zrh"
	},
}

func main() {
	str := p.Get().(string)
	fmt.Println(str)
	p.Put("kevin")
	str = p.Get().(string)
	fmt.Println(str)
}
