package authx

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JwtX struct {
	signKey      string
	signinMethod jwt.SigningMethod
}

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type Option func(s *JwtX)

func NewJwtX(opts ...Option) *JwtX {
	x := &JwtX{
		signKey:      "defaultSignKey",
		signinMethod: jwt.SigningMethodHS512,
	}
	for _, opt := range opts {
		opt(x)
	}
	return x
}

func (x *JwtX) GenerateToken(claim jwt.Claims) (string, error) {
	signKey := []byte(x.signKey)

	token := jwt.NewWithClaims(x.signinMethod, claim)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (x *JwtX) ParseToken(tokenString string, result jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, result, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(x.signKey), nil
	})

	if err != nil || !token.Valid {
		if strings.Contains(err.Error(), "expired") {
			return ErrExpiredToken
		}
		return ErrInvalidToken
	}

	return nil
}

func WithSignKey(key string) func(s *JwtX) {
	return func(s *JwtX) {
		s.signKey = key
	}
}

func WithSigingMethod(method jwt.SigningMethod) func(s *JwtX) {
	return func(s *JwtX) {
		s.signinMethod = method
	}
}
