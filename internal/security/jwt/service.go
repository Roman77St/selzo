package jwt

import (
	"errors"
	"time"

	"github.com/Roman77St/salzo/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

type Service struct {
	secret []byte

	accessTTL  time.Duration
	refreshTTL time.Duration
}

func New(secret string, accessTTL time.Duration, refreshTTL time.Duration) *Service {
	return &Service{
		secret: []byte(secret),
		accessTTL:    accessTTL,
		refreshTTL: refreshTTL,
	}
}

type TokenType string

const (
	AccessToken TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   user.Role `json:"role"`

	Type TokenType `json:"type"`

	jwt.RegisteredClaims
}

func (s *Service) GenerateAccessToken(userID uuid.UUID, role user.Role) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		Type: AccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secret)
}

func (s *Service) GenerateRefreshToken(
	userID uuid.UUID,
	role user.Role,
) (string,error) {

	claims := Claims{
		UserID: userID,
		Role: role,
		Type: RefreshToken,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(s.refreshTTL),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}


	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(s.secret)
}

func (s *Service) Parse(
	tokenString string,
) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			return s.secret, nil
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired

		default:
			return nil, ErrInvalidToken
		}
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
