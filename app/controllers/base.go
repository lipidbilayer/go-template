package controllers

import (
	"car_pool/app/core"

	"github.com/revel/revel"
)

type BaseController struct {
	*revel.Controller
	services *core.AppServices
}

func (b *BaseController) Set() revel.Result {
	b.services = core.Services
	return nil
}
