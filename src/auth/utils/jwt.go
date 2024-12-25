package utils

import (
	"Hackathon-Management-System/src/graph/model"
	configuration "Hackathon-Management-System/src/internal/config"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(user *model.User) (string, error) {

	config := configuration.NewConfig()
	expirationDurationInHours, err := strconv.Atoi(config.JWTConfig.JWT_EXPIRATION_IN_HOURS)
	if err != nil {
		return "", fmt.Errorf("error while converting string to int: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(expirationDurationInHours)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWTConfig.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (string, error) {
	config := configuration.NewConfig()

	// Validate token
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method "+" : %v", token.Header["alg"])
			}

			return []byte(config.JWTConfig.JWT_SECRET_KEY), nil
		},
	)
	if err != nil {
		return "", err
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	// Validate claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = jwt.ErrTokenInvalidClaims
		fmt.Println("Errow while claims :: ", err)
		return "", err
	}

	// Validate token expiry
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", fmt.Errorf("Error while getting expiry time from claims")
	}

	expiryTime := time.Unix(int64(exp), 0)
	if time.Now().After(expiryTime) {

		return "", fmt.Errorf("Token has expired")
	}

	// Validate User Id in the claims
	userID, ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("Error while getting user id from claims")
	}

	return userID, nil
}
