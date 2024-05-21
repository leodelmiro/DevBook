package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretkKey)
}

func ValidateToken(r *http.Request) error {
	tokenString := getToken(r)
	token, validateTokenError := jwt.Parse(tokenString, getKeyCheck)

	if validateTokenError != nil {
		return validateTokenError
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, extractUserIdError := jwt.Parse(tokenString, getKeyCheck)
	if extractUserIdError != nil {
		return 0, extractUserIdError
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, extractUserIdError := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if extractUserIdError != nil {
			return 0, extractUserIdError
		}
		
		return userID, nil
	}

	return 0, errors.New("invalid token")
}

func getKeyCheck(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Signed Key unexprected! %v", token.Header["alg"])
	}

	return config.SecretkKey, nil
}
