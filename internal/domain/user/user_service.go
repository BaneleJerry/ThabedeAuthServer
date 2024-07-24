package user

import (
	"context"

	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/database"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/repository"
	"github.com/google/uuid"
)

type UserService interface {
	GetUser(ctx context.Context, id uuid.UUID) (database.User, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
	GetUserByUsername(ctx context.Context, username string) (database.User, error)
	CreateUser(ctx context.Context, id uuid.UUID, username, passwordHash, email string) (database.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UpdateUser(ctx context.Context, id uuid.UUID, username, passwordHash, email string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(ctx context.Context, id uuid.UUID) (database.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (database.User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}

func (s *userService) CreateUser(ctx context.Context, id uuid.UUID, username, passwordHash, email string) (database.User, error) {
	params := database.CreateUserParams{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		Email:        email,
	}
	return s.repo.CreateUser(ctx, params)
}

func (s *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteUser(ctx, id)
}

// UpdateUser updates a user in the database with the provided ID, username, password hash, and email.
//
// ctx: The context for the operation.
// id: The unique identifier of the user to update.
// username: The new username for the user.
// passwordHash: The new password hash for the user.
// email: The new email for the user.
//
// Returns:
// An error if the operation fails, nil otherwise.
func (s *userService) UpdateUser(ctx context.Context, id uuid.UUID, username, passwordHash, email string) error {
	params := database.UpdateUserParams{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		Email:        email,
	}
	return s.repo.UpdateUser(ctx, params)
}
