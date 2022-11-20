package application

import (
	"github.com/jaimera/poc-services/domain/contract"
	"github.com/jaimera/poc-services/domain/service/portservice"
)

// AppService is going to hold references to all services
type AppService struct {
	app         *App
	portService contract.PortService
}

func NewAppService(app *App) *AppService {
	return &AppService{
		app: app,
	}
}

func (svc AppService) Port() contract.PortService {
	return portservice.NewPortService(svc.app.DataManager, svc.app.Logger)
}
