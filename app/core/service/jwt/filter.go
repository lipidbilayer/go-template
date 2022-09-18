package jwt

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/revel/revel"
)

func GetAuthToken(req *revel.Request) string {
	authToken := req.Header.Get("Authorization")

	if len(authToken) > 7 { // char count "Bearer " ==> 7
		return authToken[7:]
	}

	return ""
}

func AuthFilter(c *revel.Controller, fc []revel.Filter) {
	token, err := ParseToken(GetAuthToken(c.Request))
	if err == nil && token.Valid && !IsInBlocklist(GetAuthToken(c.Request)) {
		c.Args[TOKEN_CLAIMS_KEY] = token.Claims
		fc[0](c, fc[1:]) // everything looks good, move on
	} else {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				revel.AppLog.Error("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				revel.AppLog.Error("Timing is everything, Token is either expired or not active yet")
			} else {
				revel.AppLog.Error("Couldn't handle this token: %v", err)
			}
		} else {
			revel.AppLog.Error("Couldn't handle this token: %v", err)
		}

		c.Response.Status = http.StatusUnauthorized
		c.Response.Out.Header().Add("WWW-Authenticate", Realm)
		c.Result = c.RenderJSON(map[string]string{
			"id":      "unauthorized",
			"message": "Invalid or token is not provided",
		})

		return
	}

	c.Response.Status = http.StatusUnauthorized
	c.Result = c.RenderJSON(map[string]string{
		"id":      "unauthorized",
		"message": "Invalid or token is not provided",
	})
	return
	// fc[0](c, fc[1:]) //not applying JWT auth filter due to anonymous path
}
