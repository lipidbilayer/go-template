package controllers

import (
	"car_pool/app/models"
	"net/http"

	"car_pool/app/core/service/jwt"

	"github.com/revel/revel"
)

type Auth struct {
	BaseController
}

func (c Auth) Login() revel.Result {
	user, err := c.parseUserInfo()
	if err != nil {
		revel.AppLog.Errorf("Unable to read user info %s", err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]string{
			"id":      "bad_request",
			"message": "Unable to read user info",
		})
	}

	if user, err := c.services.Database.GetUser(user); err == nil {
		token, err := jwt.GenerateToken(user.ID, user.Username)
		if err != nil {
			c.Response.Status = http.StatusInternalServerError
			return c.RenderJSON(map[string]string{
				"id":      "server_error",
				"message": "Unable to generate token",
			})
		}

		return c.RenderJSON(map[string]string{
			"token": token,
		})
	}

	c.Response.Status = http.StatusUnauthorized
	// c.Response.Out.Header().Set("WWW-Authenticate", jwt.Realm)

	return c.RenderJSON(map[string]string{
		"id":      "unauthorized",
		"message": "Invalid credentials",
	})
}

func (c *Auth) parseUserInfo() (*models.User, error) {
	user := &models.User{}
	err := c.Params.BindJSON(user)
	return user, err
}
