package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth/auth/reqandresp"
)

type UserAuth struct {
	ID             string
	IP             string
	TokenType      string
	ExpireDatetime time.Time
	PairToken      string
}
type CallbackAuthentication func(string, string) string
type CallbackAuthorization func(string, string) bool

var mapAccessToken map[string]UserAuth
var mapRefreshToken map[string]UserAuth

func init() {
	mapAccessToken = make(map[string]UserAuth)
	mapRefreshToken = make(map[string]UserAuth)
}

func Authorization(e *gin.RouterGroup, cb CallbackAuthorization) {
	e.Use(func(c *gin.Context) {

		user, err := userAuth(c)
		if err != nil {
			reqandresp.Error(c, 401, err.Error())
			return
		}

		if !cb(c.Request.RequestURI, user.ID) {
			reqandresp.Unathorized(c)
			return
		}
	})
}

func Authentication(c *gin.Context) {
	user, err := userAuth(c)
	if err != nil {
		reqandresp.Error(c, 401, err.Error())
		return
	}
	if user.TokenType != reqandresp.TokenTypeResp {
		reqandresp.InvalidToken(c)
		return
	}

	if user.ExpireDatetime.Before(time.Now()) {
		reqandresp.TokenExpired(c)
		return
	}
}

func RouteToken(e *gin.Engine, cb CallbackAuthentication) {

	e.POST("/token", func(c *gin.Context) {

		grantType := c.Query("grant_type")
		//clientId := c.Query("client_id")

		switch grantType {
		case "password":
			passwordGT(c, cb)
		case "refresh_token":
			refreshTokenGT(c)
		default:
			invalidGT(c)
		}

		fmt.Println(mapAccessToken)
		fmt.Println(mapRefreshToken)
	})
}

func passwordGT(c *gin.Context, cb CallbackAuthentication) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		reqandresp.InvalidParams(c)
		return
	}

	if id := cb(username, password); id != "" {
		accessToken := reqandresp.GenerateToken()
		refreshToken := reqandresp.GenerateToken()

		mapAccessToken[accessToken] = UserAuth{
			id,
			c.ClientIP(),
			reqandresp.TokenTypeResp,
			time.Now().Add(24 * time.Hour), //1 dia
			refreshToken,
		}
		mapRefreshToken[refreshToken] = UserAuth{
			id,
			c.ClientIP(),
			reqandresp.TokenTypeResp,
			time.Now().Add(8760 * time.Hour), //1 ano
			accessToken,
		}

		reqandresp.Success(c, accessToken, refreshToken)
	} else {
		reqandresp.FailPassword(c)
	}
}

func refreshTokenGT(c *gin.Context) {

	refreshToken := c.Query("refresh_token")

	if refreshToken == "" {
		reqandresp.InvalidParams(c)
		return
	}

	userAuth, ok := mapRefreshToken[refreshToken]

	if ok {
		delete(mapRefreshToken, refreshToken)
		delete(mapAccessToken, userAuth.PairToken)

		accessT := reqandresp.GenerateToken()
		refreshT := reqandresp.GenerateToken()

		mapAccessToken[accessT] = UserAuth{
			userAuth.ID,
			c.ClientIP(),
			reqandresp.TokenTypeResp,
			time.Now().Add(24 * time.Hour), //1 dia
			refreshT,
		}

		mapRefreshToken[refreshT] = UserAuth{
			userAuth.ID,
			c.ClientIP(),
			reqandresp.TokenTypeResp,
			time.Now().Add(8760 * time.Hour), //1 ano
			accessT,
		}

		reqandresp.Success(c, accessT, refreshT)
	} else {
		reqandresp.InvalidToken(c)
	}
}

func invalidGT(c *gin.Context) {
	reqandresp.InvalidParams(c)
}

func userAuth(c *gin.Context) (*UserAuth, error) {

	authorization := c.Request.Header.Get("authorization")
	autho := strings.Split(authorization, " ")

	if authorization == "" || len(autho) != 2 {
		return nil, errors.New("invalid params")
	}
	//tokenType := autho[0]
	token := autho[1]
	user, ok := mapAccessToken[token]
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &user, nil
}
