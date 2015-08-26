package goauth

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth/authentication"
)

func Init(router *gin.Engine, cb authentication.Callback) *gin.RouterGroup {

	authentication.RouteToken(router, cb)

	aut := router.Group("/")

	aut.Use(authentication.Interceptor)

	return aut
}
