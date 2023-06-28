package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTPayload struct {
	UserId    string
	UserRoles []string
}

type TokenResult struct {
	Token  string
	Expiry int64
}

func GenerateToken(secret string, expMinutes uint, payload JWTPayload) (*TokenResult, error) {
	expiry := time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.UserId,
		"aud": payload.UserRoles,
		"exp": expiry,
	})

	s, err := t.SignedString([]byte(secret))
	if err != nil {
		return &TokenResult{}, err
	}

	return &TokenResult{Token: s, Expiry: expiry}, nil
}

func ValidateToken(secret string, token string) (*JWTPayload, error) {
	errMessage := errors.New("invalid or expired JWT")
	payload := &JWTPayload{}

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !parsed.Valid {
		return payload, errMessage
	}

	userId, err := parsed.Claims.GetSubject()
	if err != nil || userId == "" {
		return payload, errMessage
	}

	userRoles, err := parsed.Claims.GetAudience()
	if err != nil {
		return payload, errMessage
	}

	payload.UserId = userId
	payload.UserRoles = userRoles

	return payload, nil
}
