package service

import (
	"IM-Service/src/services/gateway/dao"
	"fmt"

	"gitee.com/mimis/golang-tool/net"
	"gitee.com/mimis/golang-tool/net/clientConn"
)

// 服务器统一处理客户端消息的函数(暂时放在这里，这个后面会挪到lobby服务里面去)
func HandlerRespone(reqClient *clientConn.ClientMsg) (*clientConn.ClientMsg, error) {
	fmt.Printf("client send tag:%v message:%s\n", reqClient.Tag, reqClient.Msg)
	return &clientConn.ClientMsg{
		Tag: 1,
		Msg: []byte("成功返回"),
	}, nil
}

var S *Service

type Service struct {
	// 大厅服务
	Dao *dao.Dao
}

func Init(addr string) *Service {

	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	// 服务连接etcd等服务
	s := net.InitServer(addr, "tcp", HandlerRespone)
	s.Listen()

	S = &Service{
		Dao: d,
	}
	return S
}
