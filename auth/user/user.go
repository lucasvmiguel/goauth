package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type UserAuth struct {
	ID             string
	IP             string
	TokenType      string
	ExpireDatetime time.Time
	PairToken      string
}

func CreateUserAuth(c *gin.Context, mapToken map[string]UserAuth) (*UserAuth, error) {

	authorization := c.Request.Header.Get("authorization")
	auth := strings.Split(authorization, " ")

	if authorization == "" || len(auth) != 2 {
		return nil, errors.New("invalid params")
	}

	//tokenType := autho[0]
	token := auth[1]

	user, ok := mapToken[token]

	if !ok {
		return nil, errors.New("invalid token")
	}

	return &user, nil
}
