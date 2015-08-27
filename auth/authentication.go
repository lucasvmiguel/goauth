package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth/auth/responses"
)

//CallbackAuthentication gives the power to use your logic to authenticate
//first parameter is the username
//second parameter is the password
//third parameter is the client_id
type CallbackAuthentication func(string, string, string) string

func RouteToken(e *gin.Engine, cb CallbackAuthentication) {

	e.POST("/token", func(c *gin.Context) {

		switch c.Query("grant_type") {
		case "password":
			passwordGT(c, cb)
		case "refresh_token":
			refreshTokenGT(c)
		case "destroy_token":
			destroyTokenGT(c)
		default:
			invalidGT(c)
		}
	})
}

func Authentication(c *gin.Context) {

	//TokenType underscore
	token, _, err := parseAuthorization(c.Query("authorization"))

	// fmt.Println(mapAccessToken)
	// fmt.Println(mapRefreshToken)

	if err != nil {
		responses.Error(c, 401, err.Error())
		return
	}

	user, err2 := resource.UserByAccessT(token)

	if err2 != nil {
		responses.Error(c, 401, err2.Error())
		return
	}

	if user.TokenType != responses.TokenTypeResp {
		responses.InvalidToken(c)
		return
	}

	if user.ExpireAccessDatetime.Before(time.Now()) {
		responses.TokenExpired(c)
		return
	}
}

func passwordGT(c *gin.Context, cb CallbackAuthentication) {
	username := c.Query("username")
	password := c.Query("password")
	clientID := c.Query("client_id")

	if username == "" || password == "" || clientID == "" {
		responses.InvalidParams(c)
		return
	}

	if ID := cb(username, password, clientID); ID != "" {

		user, _ := resource.InsertUser(ID, c.ClientIP())
		responses.Success(c, user.AccessToken, user.RefreshToken)
	}
	responses.FailPassword(c)
}

func refreshTokenGT(c *gin.Context) {

	refreshT := c.Query("refresh_token")

	if refreshT == "" {
		responses.InvalidParams(c)
		return
	}

	user, err := resource.UserByRefreshT(refreshT)

	if err != nil {
		responses.InvalidToken(c)
		return
	}

	if user.ExpireRefreshDatetime.Before(time.Now()) {
		responses.TokenExpired(c)
		return
	}

	err = resource.RemoveByToken(refreshT)

	if err != nil {
		responses.Error(c, 401, err.Error())
	}

	user, err = resource.InsertUser(user.ID, c.ClientIP())

	if err != nil {
		responses.Error(c, 401, err.Error())
	}
	responses.Success(c, user.AccessToken, user.RefreshToken)
}

func destroyTokenGT(c *gin.Context) {
	accessT := c.Query("access_token")
	refreshT := c.Query("refresh_token")
	clientID := c.Query("client_id")

	if accessT == "" || refreshT == "" || clientID == "" {
		responses.InvalidParams(c)
		return
	}

	if err := resource.RemoveByToken(accessT); err != nil {
		responses.Error(c*gin.Context, 401, err.Error())
	}

	responses.RemovedToken(c)
}

func invalidGT(c *gin.Context) {
	responses.InvalidParams(c)
}
