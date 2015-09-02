package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	er "github.com/lucasvmiguel/goauth/auth/errors"
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
	//SuccessProto Proto field
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

//Error send message with a error where you can define status http and message to send
func Error(c *gin.Context, err error) {

	var msgError string
	var statusError int

	switch err {
	case er.ErrInvalidParameters:
		msgError = er.ErrInvalidParameters.Error()
		statusError = 400
	case er.ErrUnauthorized:
		msgError = er.ErrUnauthorized.Error()
		statusError = 401
	case er.ErrUndefinedToken:
		msgError = er.ErrUndefinedToken.Error()
		statusError = 406
	case er.ErrTokenExpired:
		msgError = er.ErrTokenExpired.Error()
		statusError = 406
	case er.ErrValidateToken:
		msgError = er.ErrValidateToken.Error()
		statusError = 406
	case er.ErrIDRepeated:
		msgError = er.ErrIDRepeated.Error()
		statusError = 406
	default:
		msgError = er.ErrUnknown.Error()
		statusError = 406
	}

	c.JSON(statusError, gin.H{
		ErrorProto: msgError,
	})
	c.Abort()
}
