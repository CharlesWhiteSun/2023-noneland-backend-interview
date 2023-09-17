package lib

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Email string
}

type ClaimsOptions struct {
	jwt.StandardClaims
	Email string
}

type IClaimsOption interface {
	apply(*ClaimsOptions)
}

type FuncClaimsOption struct {
	f func(*ClaimsOptions)
}

func (fo FuncClaimsOption) apply(option *ClaimsOptions) {
	fo.f(option)
}

func ClaimsWithEmail(email string) IClaimsOption {
	return FuncClaimsOption{
		f: func(option *ClaimsOptions) {
			option.Email = email
		},
	}
}

func ClaimsWithIssuer(issuer string) IClaimsOption {
	return FuncClaimsOption{
		f: func(option *ClaimsOptions) {
			option.Issuer = issuer
		},
	}
}

// NewJwtObj 新增 JWT 物件
// examples:
// - NewJwtObj()
// - NewJwtObj(ClaimsWithEmail("8888@gmail.com"), ClaimsWithIssuer("apiProject"))
func NewJwtObj(id string, opts ...IClaimsOption) (*Claims, error) {
	if id == "" {
		return nil, errors.New("jwt id cannot be empty")
	}

	now := time.Now()
	init := new(ClaimsOptions)
	init.Id = id
	init.IssuedAt = now.Unix()
	init.ExpiresAt = now.Add(30 * 24 * time.Hour).Unix()

	for _, opt := range opts {
		opt.apply(init)
	}

	claim := new(Claims)
	claim.Id = init.Id
	claim.Email = init.Email
	claim.Issuer = init.Issuer
	claim.IssuedAt = init.IssuedAt
	claim.ExpiresAt = init.ExpiresAt
	return claim, nil
}

var jwtKey = []byte("FDr1VjVQiSiybYJrQZNt8Vfd7bFEsKP6vNX1brOSiWl0mAIVCxJiR4/T3zpAlBKc2/9Lw2ac4IwMElGZkssfj3dqwa7CQC7IIB+nVxiM1c9yfowAZw4WQJ86RCUTXaXvRX8JoNYlgXcRrK3BK0E/fKCOY1+izInW3abf0jEeN40HJLkXG6MZnYdhzLnPgLL/TnIFTTAbbItxqWBtkz6FkZTG+dkDSXN7xNUxlg==")

// CreateToken 製作 JWT Token String
func CreateToken(claim Claims) (string, error) {
	if claim.Id == "" {
		return "", errors.New("jwt id cannot be empty")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("jwt SignedString error: " + err.Error())
	}
	return tokenString, nil
}

// ValidateToken 驗證 JWT Token String
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, errors.New("jwt token invalid")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("jwt token invalid")
}
