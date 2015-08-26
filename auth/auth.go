package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth/auth/responses"
	"github.com/lucasvmiguel/goauth/auth/user"
	"github.com/lucasvmiguel/goauth/token"
)

type CallbackAuthentication func(string, string) string
type CallbackAuthorization func(string, string) bool

var mapAccessToken map[string]user.UserAuth
var mapRefreshToken map[string]user.UserAuth

func init() {
	mapAccessToken = make(map[string]user.UserAuth)
	mapRefreshToken = make(map[string]user.UserAuth)
}

func Authorization(e *gin.RouterGroup, cb CallbackAuthorization) {
	e.Use(func(c *gin.Context) {

		user, err := user.CreateUserAuth(c, mapAccessToken)
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

func Authentication(c *gin.Context) {
	user, err := user.CreateUserAuth(c, mapAccessToken)
	if err != nil {
		responses.Error(c, 401, err.Error())
		return
	}
	if user.TokenType != responses.TokenTypeResp {
		responses.InvalidToken(c)
		return
	}

	if user.ExpireDatetime.Before(time.Now()) {
		responses.TokenExpired(c)
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
	})
}

func passwordGT(c *gin.Context, cb CallbackAuthentication) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		responses.InvalidParams(c)
		return
	}

	if id := cb(username, password); id != "" {
		accessToken := token.Generate()
		refreshToken := token.Generate()

		mapAccessToken[accessToken] = user.UserAuth{
			id,
			c.ClientIP(),
			responses.TokenTypeResp,
			time.Now().Add(24 * time.Hour), //1 dia
			refreshToken,
		}
		mapRefreshToken[refreshToken] = user.UserAuth{
			id,
			c.ClientIP(),
			responses.TokenTypeResp,
			time.Now().Add(8760 * time.Hour), //1 ano
			accessToken,
		}

		responses.Success(c, accessToken, refreshToken)
	} else {
		responses.FailPassword(c)
	}
}

func refreshTokenGT(c *gin.Context) {

	refreshToken := c.Query("refresh_token")

	if refreshToken == "" {
		responses.InvalidParams(c)
		return
	}

	userAuth, ok := mapRefreshToken[refreshToken]

	if ok {
		delete(mapRefreshToken, refreshToken)
		delete(mapAccessToken, userAuth.PairToken)

		accessT := token.Generate()
		refreshT := token.Generate()

		mapAccessToken[accessT] = user.UserAuth{
			userAuth.ID,
			c.ClientIP(),
			responses.TokenTypeResp,
			time.Now().Add(24 * time.Hour), //1 dia
			refreshT,
		}

		mapRefreshToken[refreshT] = user.UserAuth{
			userAuth.ID,
			c.ClientIP(),
			responses.TokenTypeResp,
			time.Now().Add(8760 * time.Hour), //1 ano
			accessT,
		}

		responses.Success(c, accessT, refreshT)
	} else {
		responses.InvalidToken(c)
	}
}

func invalidGT(c *gin.Context) {
	responses.InvalidParams(c)
}
