package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	utils "tcc-test/api/utils"
)

type AuthService struct {
	userService UserService
}

type AccessToken struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{userService: *userService}
}

func (s *AuthService) Login(username string, password string) (*AccessToken, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or/and password is required")
	}
	user, err := s.userService.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Password not match.")
	}
	token, expiredAt, err := utils.GenerateAccessToken(utils.BodyToken{Username: user.Username, ID: user.ID})
	if err != nil {
		return nil, err
	}
	accessTokenObj := AccessToken{
		Token:     *token,
		ExpiredAt: *expiredAt,
	}
	return &accessTokenObj, nil
}

func (s *AuthService) Validate(token string) (*utils.ClaimsWithUser, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	parsedToken, err := jwt.ParseWithClaims(token, &utils.ClaimsWithUser{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*utils.ClaimsWithUser)
	if !ok {
		return nil, errors.New("Access token is denined.")
	}

	return claims, nil
}
