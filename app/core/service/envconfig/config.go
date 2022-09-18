package envconfig

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/revel/revel"
)

type Base struct {
	DatabaseURL string `yaml:"DATABASE_URL" env:"PORT"`

	JWTRealmName      string `yaml:"JWT_REALM_NAME" env:"JWT_REALM_NAME"`
	JWTExpiration     int64  `yaml:"JWT_EXPIRATION" env:"JWT_EXPIRATION"`
	JWTIssuerName     string `yaml:"JWT_ISSUER_NAME" env:"JWT_ISSUER_NAME"`
	JWTPublicKeyPath  string `yaml:"JWT_PUBLIC_KEY_PATH" env:"JWT_PUBLIC_KEY_PATH"`
	JWTPrivateKeyPath string `yaml:"JWT_PRIVATE_KEY_PATH" env:"JWT_PRIVATE_KEY_PATH"`
}

func New() *Base {
	var cfg Base
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		err = cleanenv.ReadConfig("../../config.yml", &cfg)
	}
	if err != nil {
		revel.AppLog.Error("Failed to read config file")
		panic(err)
	}
	return &cfg
}

func (b *Base) GetDatabaseURL() string {
	return b.DatabaseURL
}

func (b *Base) GetJWTRealmName() string {
	return b.JWTRealmName
}

func (b *Base) GetJWTExpiration() int64 {
	return b.JWTExpiration
}

func (b *Base) GetJWTIssuerName() string {
	return b.JWTIssuerName
}

func (b *Base) GetJWTPublicKeyPath() string {
	return b.JWTPublicKeyPath
}

func (b *Base) GetJWTPrivateKeyPath() string {
	return b.JWTPrivateKeyPath
}
