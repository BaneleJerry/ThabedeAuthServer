package auth

import (
	"context"

	"github.com/BaneleJerry/ThabedeAuthServer/internal/common/models"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/database"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/repository"
	"github.com/BaneleJerry/ThabedeAuthServer/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (models.User, error) {
	databaseuser, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(databaseuser.PasswordHash), []byte(password)); err != nil {
		return models.User{}, err
	}

	domainUser := utils.ConvertToDomainModel(databaseuser)

	return domainUser, nil

}

func (s *AuthService) Register(ctx context.Context, username, email, password string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	databaseUser, err := s.repo.CreateUser(ctx, database.CreateUserParams{
		ID:           uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	})

	if err != nil {
		return models.User{}, err
	}
	return utils.ConvertToDomainModel(databaseUser), nil
}
