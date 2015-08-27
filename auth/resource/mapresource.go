package resource

import (
	"time"

	"github.com/lucasvmiguel/goauth/token"
)

type Map struct {
	mapAccessToken  map[string]*User
	mapRefreshToken map[string]*User
}

func (m *Map) InsertUser(ID string, IP string) (*User, error) {
	//TODO Verificar se já existe o usuário com o mesmo id cadastrado
	accessT := token.Generate()
	refreshT := token.Generate()

	user := User{
		ID:                    ID,
		IP:                    IP,
		TokenType:             "Bearer",
		ExpireAccessDatetime:  time.Now().Add(24 * time.Hour),   //1 day
		ExpireRefreshDatetime: time.Now().Add(8760 * time.Hour), //1 year
		AccessToken:           accessT,
		RefreshToken:          refreshT,
	}

	m.mapAccessToken[accessT] = &user
	m.mapRefreshToken[refreshT] = &user

	return &user, nil
}

func (m *Map) RemoveByToken(token string) error {

	if token == "" {
		return ErrInvalidParams
	}

	accessValue, accessOk := m.mapAccessToken[token]
	refreshValue, refreshOk := m.mapRefreshToken[token]

	if !accessOk && !refreshOk {
		return ErrUndefinedToken
	}

	if accessOk {
		delete(m.mapRefreshToken, accessValue.RefreshToken)
		delete(m.mapAccessToken, accessValue.AccessToken)
	}
	if refreshOk {
		delete(m.mapRefreshToken, refreshValue.RefreshToken)
		delete(m.mapAccessToken, refreshValue.AccessToken)
	}

	return nil
}

func (m *Map) UserByAccessT(token string) (*User, error) {
	if value, ok := m.mapRefreshToken[token]; ok {
		return value, nil
	}
	return nil, ErrUndefinedToken
}

func (m *Map) UserByRefreshT(token string) (*User, error) {

	if value, ok := m.mapRefreshToken[token]; ok {
		return value, nil
	}
	return nil, ErrUndefinedToken
}
