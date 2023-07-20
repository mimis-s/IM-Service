package main

import (
	_ "embed"

	"github.com/mimis-s/IM-Service/src/common/registry"
	"github.com/mimis-s/IM-Service/src/services/account"
	"github.com/mimis-s/IM-Service/src/services/chat"
	"github.com/mimis-s/IM-Service/src/services/file"
	"github.com/mimis-s/IM-Service/src/services/friends"
	"github.com/mimis-s/IM-Service/src/services/gateway"
	"github.com/mimis-s/IM-Service/src/services/home"
	"github.com/mimis-s/IM-Service/src/services/message"
	"github.com/mimis-s/IM-Service/src/services/overrall"
	"github.com/mimis-s/IM-Service/src/services/relay"
	"github.com/mimis-s/golang_tools/app"
)

func Boots() []struct {
	bootFunc func() (*app.AppOutSideInfo, error)
} {
	list := []struct {
		bootFunc func() (*app.AppOutSideInfo, error)
	}{
		{account.CreateAppInfo},
		{chat.CreateAppInfo},
		{file.CreateAppInfo},
		{friends.CreateAppInfo},
		{gateway.CreateAppInfo},
		{home.CreateAppInfo},
		{message.CreateAppInfo},
		{overrall.CreateAppInfo},
		{relay.CreateAppInfo},
	}

	return list
}

func main() {
	s := registry.NewDefRegistry()
	bootList := Boots()
	for _, appBoot := range bootList {
		info, err := appBoot.bootFunc()
		if err != nil {
			panic(err)
		}
		s.AddAppOutSide(info)
	}

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
