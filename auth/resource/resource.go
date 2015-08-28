package resource

import "time"

//User is a struct that store the user info
type User struct {
	ID                    string
	IP                    string
	TokenType             string
	ExpireAccessDatetime  time.Time
	ExpireRefreshDatetime time.Time
	AccessToken           string
	RefreshToken          string
	Object                interface{}
}
