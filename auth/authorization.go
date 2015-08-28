package auth

import (
	"github.com/gin-gonic/gin"
	er "github.com/lucasvmiguel/goauth/auth/errors"
	resp "github.com/lucasvmiguel/goauth/auth/responses"
)

//CallbackAuthorization gives the power to use your logic to authorize
//first parameter is the route path
//second parameter is the user id
type CallbackAuthorization func(string, string) bool

//Authorization catch all requests
func Authorization(e *gin.RouterGroup, cb CallbackAuthorization) {

	e.Use(func(c *gin.Context) {

		//TokenType underscore
		_, token, err := parseAuthorization(c.Request.Header.Get("authorization"))
		if err != nil {
			resp.Error(c, err)
			return
		}

		user, err := provider.UserByAccessT(token)
		if err != nil {
			resp.Error(c, err)
			return
		}

		if !cb(c.Request.RequestURI, user.ID) {
			resp.Error(c, er.ErrUnauthorized)
			return
		}
	})
}
