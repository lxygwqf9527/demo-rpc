# Hello World Grpc

```sh
cd grpc/simple/server
# 生成service pb编译文件
protoc -I=. --go_out=. --go_opt=module="github.com/lxygwqf9527/demo-rpc/grpc/simple/server" pb/hello.proto


# 补充rpc 接口定义protobuf的代码生成
protoc -I=. --go_out=. --go_opt=module="github.com/lxygwqf9527/demo-rpc/grpc/simple/server"  \
--go-grpc_out=.  --go-grpc_opt=module="github.com/lxygwqf9527/demo-rpc/grpc/simple/server" \
pb/hello.proto
```