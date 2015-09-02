package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	er "github.com/lucasvmiguel/goauth/auth/errors"
	resp "github.com/lucasvmiguel/goauth/auth/responses"
)

//CallbackAuthentication gives the power to use your logic to authenticate:
//first parameter is the username |
//second parameter is the password |
//third parameter is the client_id
type CallbackAuthentication func(string, string, string) (string, interface{})

//RouteToken make the route "/token" and route the requests
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

		if Debug {
			provider.Debug()
		}
	})
}

//Authentication responsible to authenticate the user
func Authentication(c *gin.Context) {
	//TokenType underscore
	_, token, err := parseAuthorization(c.Request.Header.Get("authorization"))
	if err != nil {
		resp.Error(c, err)
		provider.Debug()
		return
	}

	user, err := provider.UserByAccessT(token)
	if err != nil {
		resp.Error(c, err)
		provider.Debug()
		return
	}

	if user.ExpireAccessDatetime.Before(time.Now()) {
		resp.Error(c, er.ErrTokenExpired)
		provider.Debug()
		return
	}

	provider.Debug()
}

func passwordGT(c *gin.Context, cb CallbackAuthentication) {
	username := c.Query("username")
	password := c.Query("password")
	clientID := c.Query("client_id")

	if username == "" || password == "" || clientID == "" {
		resp.Error(c, er.ErrInvalidParameters)
		return
	}

	ID, obj := cb(username, password, clientID)
	if ID == "" {
		resp.Error(c, er.ErrValidateToken)
		return
	}

	user, err := provider.InsertUser(ID, c.ClientIP(), obj)

	if err != nil {
		resp.Error(c, err)
		return
	}

	resp.Success(c, user.AccessToken, user.RefreshToken)
}

func refreshTokenGT(c *gin.Context) {

	refreshT := c.Query("refresh_token")
	if refreshT == "" {
		resp.Error(c, er.ErrInvalidParameters)
		return
	}

	user, err := provider.UserByRefreshT(refreshT)
	if err != nil {
		resp.Error(c, er.ErrUndefinedToken)
		return
	}

	if err := provider.RemoveUser(user); err != nil {
		resp.Error(c, err)
		return
	}

	user, err = provider.InsertUser(user.ID, c.ClientIP(), user.Object)
	if err != nil {
		resp.Error(c, err)
		return
	}

	resp.Success(c, user.AccessToken, user.RefreshToken)
}

func destroyTokenGT(c *gin.Context) {
	accessT := c.Query("access_token")
	refreshT := c.Query("refresh_token")
	clientID := c.Query("client_id")

	if accessT == "" || refreshT == "" || clientID == "" {
		resp.Error(c, er.ErrInvalidParameters)
		return
	}

	user, err := provider.UserByAccessT(accessT)
	if err != nil {
		resp.Error(c, er.ErrUndefinedToken)
		return
	}

	_, err = provider.UserByRefreshT(refreshT)
	if err != nil {
		resp.Error(c, er.ErrUndefinedToken)
		return
	}

	if err := provider.RemoveUser(user); err != nil {
		resp.Error(c, err)
		return
	}

	resp.RemovedToken(c)
}

func invalidGT(c *gin.Context) {
	resp.Error(c, er.ErrInvalidParameters)
}
