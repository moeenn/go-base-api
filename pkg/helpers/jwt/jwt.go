package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token jwt.Token

type JWTPayload struct {
	UserId   string
	UserRole string
}

type TokenResult struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

func GenerateToken(secret string, expMinutes uint, payload JWTPayload) (*TokenResult, error) {
	expiry := time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": payload.UserId,
		"sub": payload.UserRole,
		"exp": expiry,
	})

	s, err := t.SignedString([]byte(secret))
	if err != nil {
		return &TokenResult{}, err
	}

	return &TokenResult{Token: s, Expiry: expiry}, nil
}

func ValidateToken(secret string, token string) (*JWTPayload, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	errMessage := errors.New("Invalid or expired JWT")

	if err != nil || !parsed.Valid {
		return &JWTPayload{}, errMessage
	}

	userId, err := parsed.Claims.GetIssuer()
	if err != nil {
		return &JWTPayload{}, errMessage
	}

	userRole, err := parsed.Claims.GetSubject()
	if err != nil {
		return &JWTPayload{}, errMessage
	}

	payload := &JWTPayload{
		UserId:   userId,
		UserRole: userRole,
	}

	return payload, nil
}
