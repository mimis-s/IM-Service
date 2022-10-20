# 网页版本的客户端实现

web收发消息都是靠js封装函数实现, 序列化使用protobuf
这个通信采用websocket

# 目录解析:
assets: 存放资源文件
controller: 调用html文件方法
model: 定义客户端所用到的结构
router: 定义engine, 绑定路径和调用html方法, 运行网页
templates: html网页模板

# 想法实现
websocket在每一次登录之后都会重连, 很烦
这里结合sharedworker和iframe我有了一个新的想法, 因为sharedworker的特性, 必须有一个html页面保存之后新增页面才能实现websocket共享, 但是这个与我点击登录之后直接在本页面登录的愿望并不一致, 但是我发现iframe这个是在原有的页面嵌套使用，这样的话
相当于一定有一个框架模板来维持websocket, 这样切换界面的时候websocket就不会重连

# 规划
web客户端现在是嵌套在服务器端, 之后功能基本实现之后会考虑独立出来