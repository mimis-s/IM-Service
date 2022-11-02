# 数据库反转工具xorm reverse
https://github.com/laixyz/reverse, 使用xorm -f xxx(yaml配置文件路径)

原版将数据库url放在yaml文件conn_str字段, 项目中修改为命令行参数输入
原版: xorm -f dbmodel/reverse/custom.yaml
改版: xorm -f dbmodel/reverse/custom.yaml --dsn 'root:dev123@tcp(localhost:3306)/im_zhangbin?charset=utf8'
