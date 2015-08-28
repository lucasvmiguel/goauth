package provider

import (
	"time"

	res "github.com/lucasvmiguel/goauth/auth/resource"
)

var (
	expireAccessDatetime  = time.Now().Add(24 * time.Hour)
	expireRefreshDatetime = time.Now().Add(8760 * time.Hour)
	tokenType             = "Bearer"
)

//Provider should be implement by any provider (map, db)
type Provider interface {
	InsertUser(string, string, interface{}) (*res.User, error)
	RemoveUserByAccessT(string) error
	RemoveUserByRefreshT(string) error
	UserByAccessT(string) (*res.User, error)
	UserByRefreshT(string) (*res.User, error)

	Debug()
}
