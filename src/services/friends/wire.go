// go:build wireinject
//go:build wireinject
// +build wireinject

package friends

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/common/registry"
	"github.com/mimis-s/IM-Service/src/services/friends/job"
	"github.com/mimis-s/IM-Service/src/services/friends/service"
	"github.com/mimis-s/golang_tools/app"
	rpcxService "github.com/mimis-s/golang_tools/rpcx/service"
)

func appInject(a *app.App) (interface{}, error) {
	panic(wire.Build(
		registry.DefaultAppRpcSet,
		service.ProviderSet,
		job.InitMQ,
		registerHandler,
	))
}

func registerHandler(sm *rpcxService.ServerManage, svcHandler *service.Service, _ *job.Job, a *app.App) (interface{}, error) {

	a.AddService("rpc", sm)

	return nil, nil
}

func createAppInfo() (*app.AppOutSideInfo, error) {
	var err error
	appInfo := app.NewAppOutSide("friends",
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
