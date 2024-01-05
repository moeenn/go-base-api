package jwthelpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const (
	JWT_EXPIRY_HOURS = 72
)

type JWTCustomClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type JWTWithExpiry struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

type JWTUser struct {
	Id   string
	Role string
}

/* create JWT claim and sign using JWT secret */
func NewJWT(id, role, secret string) (JWTWithExpiry, error) {
	exp := time.Now().Add(time.Hour * JWT_EXPIRY_HOURS)
	claims := &JWTCustomClaims{
		id,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return JWTWithExpiry{}, err
	}

	return JWTWithExpiry{
		Token:  t,
		Expiry: exp.Unix(),
	}, nil
}

/* jwt middleware requires config, this function creates that config object */
func NewJWTConfig(secret string) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(secret),
	}
}

/* extract the current (logged-in) user from request context */
func CurrentUser(c echo.Context) JWTUser {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)

	return JWTUser{
		Id:   claims.Id,
		Role: claims.Role,
	}
}
