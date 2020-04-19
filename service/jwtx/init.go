package jwtx

import "os"

var jwt_key string

func init() {
	jwt_key = os.Getenv("JWT_KEY")
}
