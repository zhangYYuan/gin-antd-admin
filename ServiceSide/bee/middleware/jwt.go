package middleware

import (
	"bee/global"
	"bee/global/response"
	"bee/model/request"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired 		error = errors.New("Token is expired")
	TokenNotVaildYet 	error = errors.New("Token is expired")
	TokenMalformed		error = errors.New("Token is expired")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.BeeConfig.JWT.SigningKey),
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error){
	//token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//val, err := token.SignedString(j.SigningKey)
	//return val, err

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("token 00", token)
	return token.SignedString(j.SigningKey)
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotVaildYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.Result(response.ERROR, gin.H{
					"reload": true,
				},"授权已过期", c)
			}
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, err.Error(), c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
