package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	//TODO: delete when deploy
	_   = godotenv.Load()
	key = os.Getenv("JWT_SECRET")
)

type AuthService interface {
	GenerateToken(Id string, IpAddress string, deviceAccess string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}
func (s *authService) GenerateToken(Id string, IpAddress string, deviceAccess string) (string, error) {
	claims := jwt.MapClaims{
		"id":            Id,
		"ip_address":    IpAddress,
		"device_access": deviceAccess,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
