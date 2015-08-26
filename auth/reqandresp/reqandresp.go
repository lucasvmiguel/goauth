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
	c.Abort()
}

func InvalidToken(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "invalid token",
	})
	c.Abort()
}

func TokenExpired(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "token expired",
	})
	c.Abort()
}

func Unathorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		ErrorProto: "Unathorized",
	})
	c.Abort()
}

func Error(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		ErrorProto: msg,
	})
	c.Abort()
}
