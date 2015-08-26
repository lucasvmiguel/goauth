package authentication

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/goauth/authentication/reqandresp"
)

type UserAuth struct {
	ID             string
	IP             string
	ExpireDatetime time.Time
	PairToken      string
}
type Callback func(string, string) string

var mapAccessToken map[string]UserAuth
var mapRefreshToken map[string]UserAuth

func init() {
	mapAccessToken = make(map[string]UserAuth)
	mapRefreshToken = make(map[string]UserAuth)
}

func Interceptor(c *gin.Context) {
	c
	authorization := c.Query("authorization")
	autho := strings.Split(authorization, " ")

	if authorization == "" || len(autho) != 2 {
		reqandresp.InvalidParams(c)
		c.Abort()
	}
	//tokenType := autho[0]
	token := autho[1]

	userAuth, ok := mapAccessToken[token]

	if !ok {
		reqandresp.InvalidToken(c)
		c.Abort()
	}

	if userAuth.ExpireDatetime.Before(time.Now()) {
		reqandresp.TokenExpired(c)
		c.Abort()
	}
}

func RouteToken(c *gin.Engine, cb Callback) {

	c.POST("/token", func(c *gin.Context) {

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

func passwordGT(c *gin.Context, cb Callback) {
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
			time.Now().Add(24 * time.Hour), //1 dia
			refreshToken,
		}
		mapRefreshToken[refreshToken] = UserAuth{
			id,
			c.ClientIP(),
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
			time.Now().Add(24 * time.Hour), //1 dia
			refreshT,
		}

		mapRefreshToken[refreshT] = UserAuth{
			userAuth.ID,
			c.ClientIP(),
			time.Now().Add(8760 * time.Hour), //1 ano
			accessT,
		}

		reqandresp.Success(c, accessT, refreshT)
		return
	} else {
		reqandresp.InvalidToken(c)
	}
}

func invalidGT(c *gin.Context) {
	reqandresp.InvalidParams(c)
}
