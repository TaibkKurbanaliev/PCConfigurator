package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"pcbuilder/internal/domain"
	"pcbuilder/internal/storage"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "wqeqwsadw131d32fd32412ss"
	signingKey = "wefwefihohu197he7he14fh"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	repository storage.Authorization
}

func NewAuthService(repository storage.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) CreateUser(ctx context.Context, user domain.User) (string, error) {
	user.RegisteredAt = time.Now()
	user.LastVisitAt = time.Now()
	user.Password = generatePasswordHash(user.Password)
	return s.repository.CreateUser(ctx, user)
}

func (s *AuthService) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repository.GetUser(ctx, username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(tokenTTL).Unix(), 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(time.Now().Unix(), 0)),
		},
		UserId: user.ID.Hex(),
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(signingKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
