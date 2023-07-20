// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package chat

import (
	"github.com/mimis-s/IM-Service/src/common/registry"
	"github.com/mimis-s/IM-Service/src/services/chat/dao"
	"github.com/mimis-s/IM-Service/src/services/chat/service"
	"github.com/mimis-s/golang_tools/app"
	service2 "github.com/mimis-s/golang_tools/rpcx/service"
)

// Injectors from wire.go:

func appInject(a *app.App) (interface{}, error) {
	serverManage, err := registry.NewRpcService()
	if err != nil {
		return nil, err
	}
	daoDao, err := dao.New()
	if err != nil {
		return nil, err
	}
	serviceService, err := service.NewServiceHandler(serverManage, daoDao)
	if err != nil {
		return nil, err
	}
	v, err := registerHandler(serverManage, serviceService, a)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// wire.go:

func registerHandler(sm *service2.ServerManage, svcHandler *service.Service, a *app.App) (interface{}, error) {

	a.AddService("rpc", sm)

	return nil, nil
}

func createAppInfo() (*app.AppOutSideInfo, error) {
	var err error
	appInfo := app.NewAppOutSide("chat",
		func(a *app.App) error {
			_, err = appInject(a)
			return err

		},
	)

	return appInfo, err
}

func CreateAppInfo() (*app.AppOutSideInfo, error) {
	appInfo, err := createAppInfo()
	if err != nil {
		panic(err)
	}
	return appInfo, err
}
