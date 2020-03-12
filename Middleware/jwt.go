package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateJwt -- create jwt Code
func CreateJwt(userID uint, email, secret string) (string, error) {

	keys := []byte(secret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// Set some claims

	claims["userID"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 730).Local()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(keys)
	return tokenString, err
}

// ValidJwt -- validasi
func ValidJwt(MyToken, secret string) (jwt.MapClaims, error) {

	keys := secret

	token, err := jwt.Parse(MyToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(keys), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err

}
