package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	UUID uuid.UUID
	ID uint
	NickName string
	AuthorityId string
	jwt.StandardClaims
}