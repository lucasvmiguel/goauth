package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	//AccessTokenProto Proto field
	AccessTokenProto = "access_token"
	//TokenTypeProto Proto field
	TokenTypeProto = "token_type"
	//ExpiresInProto Proto field
	ExpiresInProto = "expires_in"
	//RefreshTokenProto Proto field
	RefreshTokenProto = "refresh_token"
	//ErrorProto Proto field
	ErrorProto = "error"
	//Success Proto field
	SuccessProto = "success"
	//TokenTypeResp Proto field
	TokenTypeResp = "Bearer"
	//ExpiresInResp Proto field
	ExpiresInResp = 3600
)

//Success send message with access_token | token_type | expires_in | refresh_token
func Success(c *gin.Context, accessToken string, refreshToken string) {
	c.JSON(http.StatusOK, gin.H{
		AccessTokenProto:  accessToken,
		TokenTypeProto:    TokenTypeResp,
		ExpiresInProto:    ExpiresInResp,
		RefreshTokenProto: refreshToken,
	})
}

//RemovedToken send message to sinalize that token was destroyed
func RemovedToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		SuccessProto: "the token was destroyed",
	})
}

//FailPassword send message with a error
func FailPassword(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "error to validate",
	})
}

//InvalidParams send message with a error
func InvalidParams(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "invalid params",
	})
	c.Abort()
}

//InvalidToken send message with a error
func InvalidToken(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "invalid token",
	})
	c.Abort()
}

//TokenExpired send message with a error
func TokenExpired(c *gin.Context) {
	c.JSON(http.StatusNotAcceptable, gin.H{
		ErrorProto: "token expired",
	})
	c.Abort()
}

//Unathorized send message with a error
func Unathorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		ErrorProto: "Unathorized",
	})
	c.Abort()
}

//Error send message with a error where you can define status http and message to send
func Error(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		ErrorProto: msg,
	})
	c.Abort()
}
