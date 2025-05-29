package jwt

import (
	"snowApp/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Manager struct {
	secretKey string
}

func NewManager(secretKey string) *Manager {
	return &Manager{
		secretKey: secretKey,
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (m *Manager) Generate(userID string, expiry time.Duration) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *Manager) Validate(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, model.ErrTokenInvalid
	}

	return claims, nil
}
