package service

import (
	"fmt"

	"gitee.com/mimis/golang-tool/net"
	"gitee.com/mimis/golang-tool/net/clientConn"
)

// 服务器统一处理客户端消息的函数
func HandlerRespone(reqClient *clientConn.ClientMsg) (*clientConn.ClientMsg, error) {
	fmt.Printf("client send tag:%v message:%s\n", reqClient.Tag, reqClient.Msg)
	return &clientConn.ClientMsg{
		Tag: 1,
		Msg: []byte("成功返回"),
	}, nil
}

func Init(addr string) {
	s := net.InitServer("localhost:8888", "tcp", HandlerRespone)
	s.Listen()
}
