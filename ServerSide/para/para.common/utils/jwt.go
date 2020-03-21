package utils

import (
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}


func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add( 3 * time.Hour)
	fmt.Println(expireTime)
	fmt.Println(jwt.SigningMethodHS256)

	clasims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jx",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, clasims)
	
	token, err := tokenClaims.SignedString(tokenClaims)
	
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
