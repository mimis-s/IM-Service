package service

import (
	chat_service "github.com/mimis-s/IM-Service/src/services/chat/service"
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/account"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/chat"
)

var S *Service

type Service struct {
	// 主界面大厅服务,主要用于处理,分发客户端的消息
	Dao  *dao.Dao
	Chat *chat_service.Service
}

func Init() *Service {
	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	S = &Service{
		Dao:  d,
		Chat: new(chat_service.Service),
	}
	return S
}
