// go:build wireinject
//go:build wireinject
// +build wireinject

package file

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/services/file/service"
	"github.com/mimis-s/golang_tools/app"
)

func appInject(a *app.App) (interface{}, error) {
	panic(wire.Build(
		service.ProviderSet,
		registerHandler,
	))
}

func registerHandler(svcHandler *service.Service, app *app.App) (interface{}, error) {

	return nil, nil
}

func createAppInfo() (*app.AppOutSideInfo, error) {
	var err error
	appInfo := app.NewAppOutSide("file",
		func(app *app.App) error {
			_, err = appInject(app)
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
