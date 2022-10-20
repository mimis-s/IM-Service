package service

import (
	"github.com/mimis-s/IM-Service/src/services/chat/dao"
)

var S *Service

type Service struct {
	Dao *dao.Dao
}

func Init() *Service {
	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	S = &Service{
		Dao: d,
	}
	return S
}
