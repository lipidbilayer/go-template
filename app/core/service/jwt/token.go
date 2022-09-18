package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthClaim struct {
	LocationID int64 `json:"location_id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int64, username string) (string, error) {
	authClaim := AuthClaim{}
	authClaim.RegisteredClaims = jwt.RegisteredClaims{
		Subject:   username,
		ID:        fmt.Sprint(id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration * time.Minute)),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	if isIssuerExists {
		authClaim.Issuer = issuer
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, authClaim)

	tokenString, err := token.SignedString(privateKey)
	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		return token, nil
	}
	return nil, jwt.NewValidationError("Claim not valid", 100)
}
