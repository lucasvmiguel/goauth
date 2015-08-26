package reqandresp

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AccessTokenProto  = "access_token"
	TokenTypeProto    = "token_type"
	ExpiresInProto    = "expires_in"
	RefreshTokenProto = "refresh_token"
	ErrorProto        = "error"

	TokenTypeResp = "Bearer"
	ExpiresInResp = 3600
	TokenSize     = 20
)

func GenerateToken() string {

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, TokenSize)
	for i := 0; i < TokenSize; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func Success(c *gin.Context, accessToken string, refreshToken string) {
	c.JSON(http.StatusOK, gin.H{
		AccessTokenProto:  accessToken,
		TokenTypeProto:    TokenTypeResp,
		ExpiresInProto:    ExpiresInResp,
		RefreshTokenProto: refreshToken,
	})
}

func FailPassword(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "error to validate",
	})
}

func InvalidParams(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "invalid params",
	})
}

func InvalidToken(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "invalid token",
	})
}

func TokenExpired(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "token expired",
	})
}
