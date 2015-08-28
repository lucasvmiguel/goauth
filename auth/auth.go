package auth

import (
	"strings"

	er "github.com/lucasvmiguel/goauth/auth/errors"
	prov "github.com/lucasvmiguel/goauth/auth/provider"
)

var (
	provider prov.Provider
	Debug    bool
)

func init() {
	//TODO aqui poderão ir outros tipos de Providers no futuro
	provider = prov.NewMap()
}

func parseAuthorization(strAuth string) (string, string, error) {

	strs := strings.Split(strAuth, " ")

	if len(strs) != 2 {
		return "", "", er.ErrInvalidParameters
	}

	return strs[0], strs[1], nil
}
