package core

import (
	"car_pool/app/core/service"
	"car_pool/app/core/service/envconfig"
	gopg "car_pool/app/core/service/go-pg"
	"car_pool/app/core/service/jwt"
)

type AppServices struct {
	Database service.DatabaseService
	Config   service.ConfigService
}

var Services *AppServices

func Init() {
	config := envconfig.New()
	database := gopg.New(config)
	jwt.Init(config)
	Services = &AppServices{Config: config, Database: database}

}

func Stop() {
	Services.Database.Stop()
}
