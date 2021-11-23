package web

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/learn/config"
)

// Create Token
// Example:
/*
	newToken, _ := createToken("Hide", "hide@gmail.com")
	log.Print("Token: ", newToken)
*/
func CreateToken(user string, email string) (string, error) {
	exp, err := time.ParseDuration(config.Config.TokenExp)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user,
		"email":    email,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(exp).Unix(),
	})
	s, err := token.SignedString([]byte(config.Config.TokenSecret))
	if err != nil {
		return "", err
	}
	return s, nil
}

// Verify Token
// Example:
/*
	info, err := verifyToken(newToken)
	if err != nil {
		log.Print("Token invalid")
	} else {
		log.Print(info)
	}
*/
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("There was an error")
		}
		return []byte(config.Config.TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid == true {
		return token.Claims.(jwt.MapClaims), err
	} else {
		return nil, fmt.Errorf("Token invalid.")
	}

}