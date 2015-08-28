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

func NewMap() *Map {
	m := Map{}
	m.mapAccessT = make(map[string]*res.User)
	m.mapRefreshT = make(map[string]*res.User)
	return &m
}

func (m Map) InsertUser(ID string, IP string, obj interface{}) (*res.User, error) {

	//TODO Verificar se já existe o usuário com o mesmo id cadastrado
	accessT := token.Generate()
	refreshT := token.Generate()

	if ID == "" || IP == "" {
		return nil, er.ErrInvalidParameters
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

func (m Map) RemoveUserByAccessT(token string) error {

	if token == "" {
		return er.ErrInvalidParameters
	}

	value, ok := m.mapAccessT[token]

	if ok {
		delete(m.mapAccessT, value.AccessToken)
		return nil
	}
	return er.ErrUndefinedToken
}

func (m Map) RemoveUserByRefreshT(token string) error {

	if token == "" {
		return er.ErrInvalidParameters
	}

	value, ok := m.mapRefreshT[token]

	if ok {
		delete(m.mapRefreshT, value.RefreshToken)
		return nil
	}
	return er.ErrUndefinedToken
}

func (m Map) UserByAccessT(token string) (*res.User, error) {

	if value, ok := m.mapAccessT[token]; ok {
		return value, nil
	}
	return nil, er.ErrUndefinedToken
}

func (m Map) UserByRefreshT(token string) (*res.User, error) {

	if value, ok := m.mapRefreshT[token]; ok {
		return value, nil
	}
	return nil, er.ErrUndefinedToken
}

//Debug help debug
func (m Map) Debug() {
	fmt.Print("Access token map: ")
	fmt.Println(m.mapAccessT)
	fmt.Print("Refresh token map: ")
	fmt.Println(m.mapRefreshT)
}
