### protobuf 相关

> 编译命令: protoc --go_out=./ *.proto

* protoc 是一个编译工具
* --go_out=./ 表示以 go 语言的格式编译，生成的文件放在当前目录
* *.proto 表示编译当前目录下所有 .proto 文件

> 注：默认情况下，protobuf 不编译 proto 文件里面定义的 service 服务。要想使之编译的话，需要使用 gRPC。
> 使用 gRPC 的编译指令：
>  * protoc --go_out=plugins=grpc:./ *.proto

### 当前目录

* client/ 基于 protobuf 生成的代码实现的 grpc 客户端
* serve/ 基于 protobuf 生成的代码实现的 grpc 服务端
* pb/ 编写 protobuf 以及编译 protobuf 生成文件的目录