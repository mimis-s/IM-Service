# IM-Service

即时通讯-服务器

# golang开发须知

1.如果想用go get依赖自己的gitee/github/gitlab库
go.mod的名字必须是项目的远程路径

2.服务之间用rpcx进行通信，rpcx以protobuf插件的形式编译得到，具体用法可参考rpcx插件：https://gitee.com/mimis/protoc-gen-rpcx

3.客户端与服务器之间的信息交互使用tcp的方式，交互形式是一个req对应一个res
  聊天服务不管是群聊还是私聊都走kcp协议，既要快速又要消息完整可靠

4.服务器消息发送可以是客户端请求返回，也可以是主动调用

# 数据库反转工具xorm reverse
https://github.com/laixyz/reverse, 使用xorm -f xxx(yaml配置文件路径)

原版将数据库url放在yaml文件conn_str字段, 项目中修改为命令行参数输入
原版: xorm -f dbmodel/reverse/custom.yaml
改版: xorm -f dbmodel/reverse/custom.yaml --dsn 'root:dev123@tcp(localhost:3306)/im_zhangbin?charset=utf8'

# 数据库结构生成
tools/sql_tools/db_sync.go文件
命令行: go run db_sync.go -d im_zhangbin -u root -p dev123

# 需要解决的困难
html iframe切换不同界面是无状态的, 也就是客户端要保存已经存在的数据是需要更大的代价
由于网页版客户端开发消耗了大量的时间,而且达不到预期的效果, 所以先考虑开发更友好的QT版客户端
尽量不在界面客户端开发浪费大量的时间, 主要目的是编写一个用于即时通讯的网络服务器

# 开发进度

# 服务器
- [x] 基本网络框架, 网页websocket, 客户端TCP, 各个服务之间rpcx
- [x] 数据库管理,导出,导入工具xorm, 每个服务dao的完善
- [x] gateway网关服搭建基本完成, 完成接收, 发送消息
- [x] home大厅服搭建基本完成, 能解析,编码,分发消息
- [x] 记录客户端的登录,注册信息,正常转发聊天消息
- [x] 数据库读写分离

# 客户端(网页版)-暂停开发
- [x] 使用websocket正常连接服务器, 并发送和接收消息
- [x] 登录界面搭建(UI和布局先不考虑), 能够正常登录服务器
- [x] 搭建简易主界面, 能和服务器发送消息
- [ ] 刷新页面也能保持登录
- [ ] 给好友发送消息
- [ ] 添加, 删除好友

# 客户端(桌面版)-持续开发中
- [ ] 连接服务器,能正常发送和接收消息
- [ ] 登录,注册功能实现
- [ ] 好友系统搭建,获取好友列表; 添加,删除好友
- [ ] 正常和好友聊天
- [ ] 群聊系统,群聊; 添加, 删除群
