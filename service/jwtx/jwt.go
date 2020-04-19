package jwtx

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenToken(params map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for k, v := range params {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(jwt_key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(jwt_key), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token vaild faild")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims failed")
	}
	return claims, nil
}
