package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

type Context struct {
	ID uint64
	Username string
}

func secretFunc(secret string) jwt.Keyfunc {
	return func (token *jwt.Token) (interface{} error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}
