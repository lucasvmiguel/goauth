package auth

import (
	"errors"
	"strings"

	res "github.com/lucasvmiguel/goauth/auth/resource"
)

//CallbackAuthorization gives the power to use your logic to authorize
//first parameter is the route path
//second parameter is the user id
type CallbackAuthorization func(string, string) bool

//Authorization catch all requests
func Authorization(e *gin.RouterGroup, cb CallbackAuthorization) {

	e.Use(func(c *gin.Context) {

		//TokenType underscore
		token, _, err := parseAuthorization(c.Query("authorization"))

		if err != nil {
			responses.Error(c, 401, err.Error())
			return
		}

		user, err := resource.UserByAccessT(token)

		if err != nil {
			responses.Error(c, 401, err.Error())
			return
		}

		if !cb(c.Request.RequestURI, user.ID) {
			responses.Unathorized(c)
			return
		}
	})
}
