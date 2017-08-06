package JWT

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/NiciiA/AuthRest/application/domain"
	"fmt"
)

var Signed string = "lkjaslc4Ajskdfhr3ucihjsfdh"

func CreateToken(u Domain.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   u.ID.Hex(),
		"role": u.Role,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().AddDate(4, 0, 0).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(Signed))
	return tokenString
}

func DecodeToken(t string) (jwtoken *jwt.Token, err error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Signed), nil
	})
	return token, err
}

func IsLoggedIn(t string) bool {
	token, err := DecodeToken(t)
	if err == nil && token.Valid {
		return true
	}
	return false
}

func TokenClaims(t string) jwt.MapClaims {
	token, err := DecodeToken(t)
	if err == nil && token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			return claims
		}
	}
	return nil
}