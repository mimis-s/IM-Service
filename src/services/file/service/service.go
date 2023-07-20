package service

// import (
// 	"github.com/mimis-s/IM-Service/src/common/boot_config"
// 	"github.com/mimis-s/IM-Service/src/services/file/dao"
// 	"github.com/mimis-s/golang_tools/net"
// )

// var S *Service

// type Service struct {
// 	Dao *dao.Dao
// }

// func Init(configOptions *boot_config.ConfigOptions) *Service {

// 	d, err := dao.New(configOptions)
// 	if err != nil {
// 		panic(err)
// 	}

// 	S = &Service{
// 		Dao: d,
// 	}

// 	tcpAddr := configOptions.CommandFlags.FileServiceIP + ":" + configOptions.CommandFlags.FileServicePort

// 	// 客户端连接的TCP连接
// 	tcpServer := net.InitServer(tcpAddr, "tcp", NewSession)

// 	go func() {
// 		err := tcpServer.Run()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}()

// 	return S
// }

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/services/chat/dao"
	"github.com/mimis-s/golang_tools/net"
)

var ProviderSet = wire.NewSet(dao.ProviderSet, NewServiceHandler)

var S *Service

type Service struct {
	Dao *dao.Dao
}

func NewServiceHandler(d *dao.Dao) (*Service, error) {

	S = &Service{
		Dao: d,
	}

	tcpAddr := boot_config.CustomBootFlagsData.FileServiceIP + ":" + boot_config.CustomBootFlagsData.FileServicePort

	// 客户端连接的TCP连接
	tcpServer := net.InitServer(tcpAddr, "tcp", NewSession)

	go func() {
		err := tcpServer.Run()
		if err != nil {
			panic(err)
		}
	}()

	return S, nil
}
