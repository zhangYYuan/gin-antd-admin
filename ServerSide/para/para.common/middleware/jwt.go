package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"para.common/conf"
	"para.common/models"
	"time"
)

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(conf.GlobalConfig.JWT.SigningKey)}
}

// 创建一个token
func (j *JWT) GenerateToken (user models.SysUser) (string, error) {
	expirationTime := time.Now().Add(7 * 24 *time.Hour)

	claims := &Claims{
		UserId:         1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "bubble",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//fmt.Println(user.ID, j.SigningKey)
	return token.SignedString(j.SigningKey)
}

func (j *JWT)ParseToken(tokenString string)(*jwt.Token, *Claims, error)  {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	return token, claims, err
}