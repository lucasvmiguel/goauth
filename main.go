package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth"
)

func main() {
	router := gin.Default()

	//callback to validate authentication
	a := goauth.Init(router, func(username string, password string) string {
		if username == "lucas" && password == "123456" {
			return "1"
		}
		return ""
	})

	a.GET("/admin", func(c *gin.Context) {
		c.String(200, "TESTE")
	})

	router.Run(":8080")
}
