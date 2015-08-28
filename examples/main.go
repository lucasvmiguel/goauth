package main

import (
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

	//Initialize the goauth
	//router - gin.Engine
	//ctrlAuthentication - controller used to authenticate
	//ctrlAuthorization - controller used to authorize
	a := goauth.Init(router, ctrlAuthentication, ctrlAuthorization, true)

	//create route with goauth init return if you DO NOT want to authenticate the route
	//AUTHENTICATE
	a.GET("/admin", func(c *gin.Context) {
		c.String(200, "ADMIN")
	})

	//create route with router(gin.Engine) if you DO NOT want to authenticate the route
	//NOT AUTHENTICATE
	router.GET("/teste", func(c *gin.Context) {
		c.String(200, "TESTE")
	})

	router.Run(":8082")
}

//controller authentication gives you the username, password and clientID
//you need return the id and the object user(optional) if you want authenticate
func ctrlAuthentication(username string, password string, clientID string) (string, interface{}) {
	if username == "lucas" && password == "123456" {
		//AUTHENTICATE
		return "1", nil
	}
	//NOT AUTHENTICATE
	return "", nil
}

//controller authentication gives you the path and id user
//you just need return true to authorize or false to unauthorize
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
