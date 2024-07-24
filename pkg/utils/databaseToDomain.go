package utils

import (
	"github.com/BaneleJerry/ThabedeAuthServer/internal/common/models"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/database"
)

func ConvertToDomainModel(u database.User) models.User {
	return models.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}
