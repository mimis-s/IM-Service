网页版本的客户端实现

web收发消息都是靠js封装函数实现, 序列化使用protobuf

目录解析:
assets: 存放资源文件
controller: 调用html文件方法
model: 定义客户端所用到的结构
router: 定义engine, 绑定路径和调用html方法, 运行网页
templates: html网页模板