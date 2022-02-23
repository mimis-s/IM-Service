# IM-Service

即时通讯-服务器

# golang开发须知

1.如果想用go get依赖自己的gitee/github/gitlab库
go.mod的名字必须是项目的远程路径

2.服务之间用rpcx进行通信，rpcx以protobuf插件的形式编译得到，具体用法可参考rpcx插件：https://gitee.com/mimis/protoc-gen-rpcx