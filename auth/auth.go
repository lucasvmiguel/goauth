package auth

import (
	res "github.com/lucasvmiguel/goauth/auth/resource"
	"strings"
)

var resource *res.Resourcer

func init() {
	resource = res.Resourcer(res.Map{})
}

func parseAuthorization(strAuth string) (string, string, error) {

	strs := strings.Split(strAuth, " ")

	if len(strs) != 2 {
		return "", "", ErrInvalidParameter
	}

	return strs[0], strs[1], nil
}
