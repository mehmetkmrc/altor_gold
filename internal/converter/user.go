package converter

import (
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/entity"
	"github.com/mehmetkmrc/ator_gold/internal/dto"
)

func GetUserModelToDto(userData *entity.User) *dto.GetUserResponse {
	return &dto.GetUserResponse{
		UserID:    userData.UserID,
		Name:      userData.Name,
		Surname:   userData.Surname,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}
}
