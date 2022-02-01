package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// Payload define payload body for token
type Payload struct {
	Username     string
	Phone        string
	Role         string
	RegisteredAt string
}

type Claims struct {
	Payload
	jwt.StandardClaims
}

type GetToken struct {
	RefreshToken          string
	AccessToken           string
	RefreshTokenExpiresAt time.Time
	AccessTokenExpiresAt  time.Time
}
