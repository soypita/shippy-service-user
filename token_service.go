package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("mySuperSecretKeyLol")
)

type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *User) (string, error)
}

// Decode a token string into a token object
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "go.micro.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
