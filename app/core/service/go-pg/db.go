package gopg

import (
	"car_pool/app/core/service"

	"github.com/go-pg/pg/v10"
	"github.com/revel/revel"
)

var db *pg.DB

func New(config service.ConfigService) *DatabaseGoPG {
	opt, err := pg.ParseURL(config.GetDatabaseURL())
	if err != nil {
		revel.AppLog.Errorf("failed to parse url")
		panic(err)
	}
	db = pg.Connect(opt)
	dbService := DatabaseGoPG{db}
	return &dbService
}

func (d *DatabaseGoPG) Stop() {
	if err := db.Close(); err != nil {
		revel.AppLog.Error("Failed to close the database", "error", err)
		panic(err)
	}
}
