package resource

import (
	"errors"
	"time"
)

var (
	//ErrUndefinedToken undefined token
	ErrUndefinedToken = errors.New("Undefined error")
	//ErrInvalidParams invalid params
	ErrInvalidParams = errors.New("Invalid error")
)

//User is a struct that store the user info
type User struct {
	ID                    string
	IP                    string
	TokenType             string
	ExpireAccessDatetime  time.Time
	ExpireRefreshDatetime time.Time
	AccessToken           string
	RefreshToken          string
}

//Resourcer should be implement by any resource (map, db)
type Resourcer interface {
	InsertUser(ID string, IP string) (*User, error)
	RemoveByToken(token string) error
	UserByAccessT(token string) (*User, error)
	UserByRefreshT(token string) (*User, error)
}
