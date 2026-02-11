package util

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ClaimsWithUser struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type BodyToken struct {
	ID       string
	Username string
}

// HashPassword takes a plain-text password and returns a bcrypt hash.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a plain-text password with a bcrypt hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Validate(token string) (*ClaimsWithUser, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	parsedToken, err := jwt.ParseWithClaims(token, &ClaimsWithUser{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*ClaimsWithUser)
	if !ok {
		return nil, errors.New("Access token is denined.")
	}

	return claims, nil
}

func GenerateAccessToken(user BodyToken) (*string, *time.Time, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	expSecond, err := strconv.Atoi(os.Getenv("JWT_TIME_EXPIRED_SECOND"))
	if err != nil {
		return nil, nil, err
	}

	expirationTime := time.Now().Add(time.Duration(expSecond) * time.Second)
	claims := ClaimsWithUser{
		user.Username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        user.ID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, nil, err
	}
	return &tokenString, &expirationTime, nil
}
