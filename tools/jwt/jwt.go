package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

var secretKey = []byte("secret-key")

func CreateTokenJWT(username string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"userId":   userId,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func GetUserIDFromToken(c *gin.Context) (string, error) {
	tokenString, errToken := GetTokenAuth(c)
	if errToken != nil {
		return "", errToken
	}

	log.Debug().Msg("tokenString" + tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var username string
		switch v := claims["userId"].(type) {
		case string:
			username = v
		case float64:
			username = fmt.Sprintf("%v", v)
		default:
			return "", fmt.Errorf("userId not found in token")
		}

		return username, nil
	}

	return "", errors.New("invalid token")
}

func GetTokenAuth(c *gin.Context) (string, error) {

	authHeader := c.Request.Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return "", errors.New("token not found")
	}

	authHeader = authHeader[len("Bearer "):]

	return authHeader, nil
}
