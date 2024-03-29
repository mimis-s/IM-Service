// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package file

import (
	"github.com/mimis-s/IM-Service/src/services/chat/dao"
	"github.com/mimis-s/IM-Service/src/services/file/service"
	"github.com/mimis-s/golang_tools/app"
)

// Injectors from wire.go:

func appInject(a *app.App) (interface{}, error) {
	daoDao, err := dao.New()
	if err != nil {
		return nil, err
	}
	serviceService, err := service.NewServiceHandler(daoDao)
	if err != nil {
		return nil, err
	}
	v, err := registerHandler(serviceService, a)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// wire.go:

func registerHandler(svcHandler *service.Service, app2 *app.App) (interface{}, error) {

	return nil, nil
}

func createAppInfo() (*app.AppOutSideInfo, error) {
	var err error
	appInfo := app.NewAppOutSide("file",
		func(app2 *app.App) error {
			_, err = appInject(app2)
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
