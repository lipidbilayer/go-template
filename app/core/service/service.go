package service

import "car_pool/app/models"

type DatabaseService interface {
	Stop()
	GetUser(*models.User) (*models.User, error)
}

type ConfigService interface {
	GetDatabaseURL() string
	GetJWTRealmName() string
	GetJWTExpiration() int64
	GetJWTIssuerName() string
	GetJWTPublicKeyPath() string
	GetJWTPrivateKeyPath() string
}
