package provider

import (
	er "github.com/lucasvmiguel/goauth/auth/errors"
	res "github.com/lucasvmiguel/goauth/auth/resource"
	"github.com/lucasvmiguel/goauth/token"

	"fmt"
)

//Map is a provider to maintain tokens/users
type Map struct {
	mapAccessT  map[string]*res.User
	mapRefreshT map[string]*res.User
}

//New creates a provider
func New() *Map {
	m := Map{}
	m.mapAccessT = make(map[string]*res.User)
	m.mapRefreshT = make(map[string]*res.User)
	return &m
}

//InsertUser insert user in provider(map)
func (m Map) InsertUser(ID string, IP string, obj interface{}) (*res.User, error) {

	accessT := token.Generate()
	refreshT := token.Generate()

	if ID == "" || IP == "" {
		return nil, er.ErrInvalidParameters
	}

	for _, v := range m.mapAccessT {
		if v.ID == ID {
			return nil, er.ErrIDRepeated
		}
	}

	for _, v := range m.mapRefreshT {
		if v.ID == ID {
			return nil, er.ErrIDRepeated
		}
	}

	user := res.User{
		ID:                    ID,
		IP:                    IP,
		TokenType:             tokenType,
		ExpireAccessDatetime:  expireAccessDatetime,
		ExpireRefreshDatetime: expireRefreshDatetime,
		AccessToken:           accessT,
		RefreshToken:          refreshT,
		Object:                obj,
	}

	m.mapAccessT[accessT] = &user
	m.mapRefreshT[refreshT] = &user

	return &user, nil
}

//RemoveUser remove user from provider(map)
func (m Map) RemoveUser(user *res.User) error {

	if user.AccessToken == "" || user.RefreshToken == "" {
		return er.ErrInvalidParameters
	}

	_, okAccess := m.mapAccessT[user.AccessToken]
	_, okRefresh := m.mapRefreshT[user.RefreshToken]

	if okAccess && okRefresh {
		delete(m.mapAccessT, user.AccessToken)
		delete(m.mapRefreshT, user.RefreshToken)
		return nil
	}
	return er.ErrUndefinedToken
}

//UserByAccessT get user from provider(map) by access token
func (m Map) UserByAccessT(token string) (*res.User, error) {

	if value, ok := m.mapAccessT[token]; ok {
		return value, nil
	}
	return nil, er.ErrUndefinedToken
}

//UserByRefreshT get user from provider(map) by refresh token
func (m Map) UserByRefreshT(token string) (*res.User, error) {

	if value, ok := m.mapRefreshT[token]; ok {
		return value, nil
	}
	return nil, er.ErrUndefinedToken
}

//Debug help debug
func (m Map) Debug() {
	fmt.Print("Access map token: ")
	fmt.Println(m.mapAccessT)
	fmt.Print("Refresh map token: ")
	fmt.Println(m.mapRefreshT)
}
