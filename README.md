projeto de autorização e autenticação em golang usando o flow Password credentials(OAuth2)

package main

import (
	_ "fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth"
)

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})

	//callback to validate authentication
	a := goauth.Init(router, ctrlAuthentication, ctrlAuthorization, true)

	a.GET("/profile", func(c *gin.Context) {
		roles := []string{"admin", "teste"}

		c.JSON(200, gin.H{
			"email": "lucas@gmail.com",
			"name":  "lucas",
			"roles": roles,
		})
	})

	a.GET("/admin", func(c *gin.Context) {
		c.String(200, "ADMIN")
	})

	a.GET("/teste", func(c *gin.Context) {
		c.String(200, "TESTE")
	})

	router.Run(":8082")
}

func ctrlAuthentication(username string, password string, clientID string) (string, interface{}) {
	if username == "lucas" && password == "123456" {
		return "1", nil
	}
	return "", nil
}

func ctrlAuthorization(path string, id string) bool {
	if id == "1" && path == "/profile" {
		return true
	}
	if id == "1" && path == "/admin" {
		return true
	}
	if id == "2" && path == "/teste" {
		return true
	}

	return false
}
                                                               
