package service

import (
	"time"

	"github.com/NatananPh/kiosk-machine-api/pkg/auth/repository"
	"github.com/golang-jwt/jwt/v5"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
	}
}

func (a *AuthServiceImpl) Login(username, password string) (string, error) {
	user, err := a.AuthRepository.GetAuthUser(username, password)
	if err != nil {
		return "", err
	}

	claims := &jwtCustomClaims{
		user.Username,
		user.RoleID == 1,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
