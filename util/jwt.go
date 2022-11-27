package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var hmacSampleSecret = []byte("123456789")

func SignJwt(s string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data_str": s,
		"expire":   time.Now().Add(time.Second + 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)

	return tokenString
}

func ParseJwt(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["data_str"], claims["expire"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
