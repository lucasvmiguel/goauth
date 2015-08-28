package goauth

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth/auth"
)

func Init(router *gin.Engine, cbAuthentication auth.CallbackAuthentication, cbAuthorization auth.CallbackAuthorization, debug bool) *gin.RouterGroup {

	auth.Debug = debug

	auth.RouteToken(router, cbAuthentication)

	aut := router.Group("/")

	aut.Use(auth.Authentication)

	auth.Authorization(aut, cbAuthorization)

	return aut
}
