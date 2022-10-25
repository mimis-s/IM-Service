package service

import (
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	"github.com/mimis-s/golang_tools/net"
)

var S *Service

// 现在的rpcx调用都不用,先使用本地调用
type Service struct {
	Dao *dao.Dao
}

func Init(addr, webAddr string) *Service {

	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	S = &Service{
		Dao: d,
	}

	httpServer := net.InitServer(webAddr, "http", NewSession)

	go func() {
		err := httpServer.Listen()
		if err != nil {
			panic(err)
		}
	}()

	// 客户端连接的TCP连接
	// s := net.InitServer(addr, "tcp", HandlerRespone)
	// s.Listen()

	return S
}
