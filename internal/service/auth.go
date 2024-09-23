package service

import (
	"context"
	"crypto/sha1"
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

func (s *AuthService) CreateUser(ctx context.Context, user domain.User) (int, error) {
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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
