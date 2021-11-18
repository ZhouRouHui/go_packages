package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// Login 定义接收数据的结构体
type Login struct {
	// binding:"required 修饰的字段，若接收为空值，则报错，是必需字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "zrh")
	c.String(200, fmt.Sprintf("hello %s", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "kevin")
	c.String(200, fmt.Sprintf("hello %s", name))
}

// GlobalMiddleWare 全局中间件
func GlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到 Context 的 key 中，可以通过 Get() 取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()

		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func PartMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("局部中间件开始")
		c.Set("part", "part111")
		c.Next()

		fmt.Println("局部中间件结束")
	}
}

func main() {
	// 创建路由
	// Default() 返回的实例默认使用了 2 个中间件Logger(), Recovery()
	// 如果不需要可以使用 gin.New() 来得到新的实例
	r := gin.Default()

	// --------------------------------------- 设置路由 -----------------------------------------

	// 绑定路由规则和执行的函数
	// gin.Context 封装了 request 和 response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})

	// 绑定路由规则和执行的函数
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 路由组
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	// --------------------------------------- 中间件 -----------------------------------------
	// 注册全局中间件
	r.Use(GlobalMiddleWare())
	{
		r.GET("/md", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})
	}

	// 局部中间件
	r.GET("partMD", PartMiddleWare(), func(c *gin.Context) {
		// 取值
		req, _ := c.Get("part")
		fmt.Println("part:", req)
		// 页面接收
		c.JSON(200, gin.H{"part": req})
	})

	// --------------------------------------- 获取参数 -----------------------------------------

	// 获取 uri 中的参数, Param() 方法来获取API参数
	// 例：localhost:8000/xxx/zhangsan
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(name)   // /user/zrh
		fmt.Println(action) // /get
		// 截取 action 中的 /
		action = strings.Trim(action, "/")
		fmt.Println(action)

		c.String(http.StatusOK, name+" is "+action)
	})

	// 获取 url 的 query 参数, DefaultQuery() 或 Query() 方法
	// DefaultQuery() 若参数不存在，返回默认值，Query() 若不存在，返回空串
	// 例：localhost:8080?name=kevin
	r.GET("/user", func(c *gin.Context) {
		// 指定默认值
		// http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "zrh")
		c.String(http.StatusOK, fmt.Sprintf("hello, %s", name))
	})

	// 表单参数，通过 DefaultPostForm() 和 PostForm() 方法获取，该方法默认解析的是 x-www-form-urlencoded 或 form-data 格式的参数
	// 表单传输为 post 请求，http 常见的传输格式为四种：
	// 		application/json
	// 		application/x-www-form-urlencoded
	//		application/xml
	// 		multipart/form-data
	r.POST("/form", func(c *gin.Context) {
		reqType := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, type: %s", username, password, reqType))
	})

	// 上传单个文件， FormFile() 方法获取图片内容
	// gin 文件上传与原生的 net/http 方法类似，不同在于 gin 把原生的 request 封装到 c.Request 中
	// 限制上传最大尺寸, 限制表单上传大小 8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	// 上传多个文件，MultipartForm() 方法
	r.POST("/multi-upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err = c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err))
				return
			}
		}

		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	// --------------------------------------- 解析与绑定 -----------------------------------------

	// Json 数据解析和绑定
	// ShouldBindJSON() 和 BindJSON() 方法进行绑定
	// 两者的区别：Shouldxxx 和 bindxxx 区别就是当解析错误时 bindxxx 会在 header 中添加 400 的返回信息，而 Shouldxxx 不会添加，
	// Shouldxxx 需要给客户端返回什么状态码由你自己决定
	// 客户端传参，后端接收并解析到结构体
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将 request 的 body 中的数据，自动按照 json 格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误消息
			// gin.H 封装了生成 json 数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// 表单数据解析和绑定，Bind() 和 ShouldBind() 方法
	r.POST("loginForm", func(c *gin.Context) {
		var form Login
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// URI数据解析和绑定, ShouldBindUri() 和 BindUri() 方法
	r.POST("/:user/:password", func(c *gin.Context) {
		var uri Login
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if uri.User != "root" || uri.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// --------------------------------------- 数据渲染响应 -----------------------------------------

	// json 渲染响应
	r.GET("/renderJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "render json", "status": 200})
	})

	// 结构体渲染响应
	r.GET("renderStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})

	// xml 渲染响应
	r.GET("renderXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})

	// yaml 渲染响应
	r.GET("renderYAML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})

	// protobuf 格式，谷歌开发的高效存储读取的工具
	r.GET("renderProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传 protobuf 格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	// html 渲染
	// gin 支持加载 html 模板，然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换
	// LoadHTMLGlob() 方法可以加载模板文件
	r.LoadHTMLGlob("tpl/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
	})

	// --------------------------------------- 会话控制 -----------------------------------------

	// HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出
	// Cookie就是解决HTTP协议无状态的方案之一，中文是小甜饼的意思
	// Cookie实际上就是服务器保存在浏览器上的一段信息。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
	// Cookie由服务器创建，并发送给浏览器，最终由浏览器保存
	r.GET("set_cookie", func(c *gin.Context) {
		// 获取客户端是否携带 cookie
		cookie, err := c.Cookie("test_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置 cookie
			// maxAge int，单位为秒
			// path，cookie 所在目录
			// domain string，域名
			// secure 是否只能通过 https 访问
			// httpOnly bool 是否允许别人通过 js 获取自动的 cookie
			c.SetCookie("test_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie 的值是：%s\n", cookie)
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	// --------------------------------------- 参数验证 -----------------------------------------

	// 结构体验证
	r.GET("/5lmh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})

	// 自定义验证
	// 对绑定解析到结构体上的参数，自定义验证功能
	// 比如我们要对 hobby 字段做校验，要不能为空，并且不等于 football ，类似这种需求，就无法 binding 现成的方法
	// 需要我们自己验证方法才能实现 官网示例（https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Custom_Functions）
	// 这里需要下载引入下 gopkg.in/go-playground/validator.v8

	// 3、将我们自定义的校验方法注册到 validator中
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	// 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
	//	v.RegisterValidation("NotNullAndAdmin", hobbyNotNullAndAdmin)
	//}

	r.GET("/customize", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})

	// --------------------------------------- 启动服务 -----------------------------------------
	// 监听端口，默认在 8080
	// Run("里面不指定端口号默认为8080")
	r.Run()
}

// Person 数据校验测试结构体
type Person struct {
	// 不能为空且大于 10
	Age  int    `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
	// 在参数 binding 上使用自定义的校验方法函数注册时候的名称
	Hobby    string    `form:"hobby" binding:"NotNullAndFootball"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

// 自定义数据校验方法
func hobbyNotNullAndAdmin(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if value, ok := field.Interface().(string); ok {
		// 字段不能为空，并且不等于  football
		return value != "" && !("football" == value)
	}

	return true
}

// AuthMiddleWare auth 信息校验中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端 cookie 并校验
		if cookie, err := c.Cookie("test_cookie"); err == nil {
			if cookie == "value_cookie" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth middleware failed"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
