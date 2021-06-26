package request

import uuid "github.com/satori/go.uuid"
import "github.com/dgrijalva/jwt-go"

// Custom claims structure

type CustomClaims struct {
	UUID       uuid.UUID
	ID         uint
	Username   string
	NickName   string
	BufferTime int64
	jwt.StandardClaims
}
