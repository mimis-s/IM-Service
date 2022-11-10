package service

import (
	_ "github.com/mimis-s/IM-Service/src/services/home/service/account"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/chat"
)

var S *Service

type Service struct {
}

func Init() *Service {

	S = &Service{}
	return S
}
