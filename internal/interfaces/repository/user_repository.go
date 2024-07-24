package repository

import (
	"context"

	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/database"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (database.User, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
	GetUserByUsername(ctx context.Context, username string) (database.User, error)
	CreateUser(ctx context.Context, arg database.CreateUserParams) (database.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UpdateUser(ctx context.Context, arg database.UpdateUserParams) error
}

type userRepository struct {
	queries *database.Queries
}

func NewUserRepository(queries *database.Queries) UserRepository {
	return &userRepository{queries: queries}
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (database.User, error) {
	sqlcUser, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return database.User{}, err
	}
	return sqlcUser, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	sqlcUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return database.User{}, err
	}
	return sqlcUser, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (database.User, error) {
	sqlcUser, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return database.User{}, err
	}
	return sqlcUser, nil
}

func (r *userRepository) CreateUser(ctx context.Context, arg database.CreateUserParams) (database.User, error) {
	sqlcUser, err := r.queries.CreateUser(ctx, arg)
	if err != nil {
		return database.User{}, err
	}
	return sqlcUser, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *userRepository) UpdateUser(ctx context.Context, arg database.UpdateUserParams) error {
	return r.queries.UpdateUser(ctx, arg)
}
