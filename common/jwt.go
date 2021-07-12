/*
 * @Author: laf
 * @Date: 2021-07-12 20:13:31
 * @LastEditTime: 2021-07-12 19:47:02
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\common\jwt.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package common

import (
	"time"

	"github.com/golang-jwt/jwt"
	"laf.ltd/rujie/model"
)

var jwtkey = []byte("a_secret_correct")

// Claims struct
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	Claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "laf.ltd",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken ParseToken
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})

	return token, claims, err
}
