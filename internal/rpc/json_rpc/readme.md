### jsonrpc 案例

> jsonrpc：有效解决由在 go 语言环境中启动的客户端或服务端使用 go 语言独有的 gob 数据序列化方式而引起的另一端接续数据乱码的问题。

* 客户端

```go
// 使用 jsonrpc 来创建链接
conn, err := jsonrpc.Dial("tcp", ":9001")
```

* 服务端

```go
// 使用 jsonrpc 绑定服务
jsonrpc.ServeConn(conn)
```
