package service

import (
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
)

var S *Service

type Service struct {
	// 主界面大厅服务,主要用于处理,分发客户端的消息
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
	return S
}
