package service

import (
	"context"
	"snowApp/internal/model"
	"snowApp/internal/repository"
	"snowApp/pkg/jwt"
	"time"
)

type AuthService struct {
	userRepo      repository.UserRepository
	jwtManager    *jwt.Manager
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
}

func NewAuthService(
	userRepo repository.UserRepository,
	jwtManager *jwt.Manager,
	tokenExpiry, refreshExpiry time.Duration) *AuthService {
	return &AuthService{
		userRepo:      userRepo,
		jwtManager:    jwtManager,
		tokenExpiry:   tokenExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (s *AuthService) Register(ctx context.Context, username, email, birthday, avatar_url, bio, password string) (*model.User, error) {
	_, err := s.userRepo.FindByEmail(ctx, email)
	if err == nil {
		return nil, model.ErrUserAlreadyExists
	}
	user := &model.User{
		Username:  username,
		Email:     email,
		AvatarUrl: avatar_url,
		Bio:       bio,
		Birthday:  time.Time{}, // Convert birthday string to time.Time if needed
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	_, err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*model.User, string, string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", "", model.ErrEmailNotFound
	}

	if !user.CheckPassword(password) {
		return nil, "", "", model.ErrUserNotFound
	}

	// Generate tokens
	accessToken, err := s.jwtManager.Generate(user.ID, s.tokenExpiry)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := s.jwtManager.Generate(user.ID, s.refreshExpiry)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (string, error) {
	claims, err := s.jwtManager.Validate(token)
	if err != nil {
		return "", err
	}

	// Check if user exists
	_, err = s.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		return "", model.ErrUserNotFound
	}

	return claims.UserID, nil
}
