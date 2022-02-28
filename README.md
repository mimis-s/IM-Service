# IM-Service

即时通讯-服务器

# golang开发须知

1.如果想用go get依赖自己的gitee/github/gitlab库
go.mod的名字必须是项目的远程路径

2.服务之间用rpcx进行通信，rpcx以protobuf插件的形式编译得到，具体用法可参考rpcx插件：https://gitee.com/mimis/protoc-gen-rpcx

3.客户端与服务器之间的信息交互使用tcp的方式，交互形式是一个req对应一个res
  聊天服务不管是群聊还是私聊都走kcp协议，既要快速又要消息完整可靠

4.服务器消息发送可以是客户端请求返回，也可以是主动调用