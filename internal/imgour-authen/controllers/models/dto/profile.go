package dto

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/models"
	"github.com/gin-gonic/gin"
)

type ProfileDto struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Phone string   `json:"phone"`
	Roles []string `json:"roles"`
}

func (dto ProfileDto) ParseToResponse() *gin.H {
	return &gin.H{
		"id":    dto.Id,
		"name":  dto.Name,
		"email": dto.Email,
		"phone": dto.Phone,
		"roles": dto.Roles,
	}
}

func ParseFromModel(profile *models.UserProfile) *ProfileDto {
	var roles []string
	for _, role := range profile.Roles {
		roles = append(roles, models.Role_name[role])
	}

	return &ProfileDto{
		Id:    profile.Uid,
		Name:  profile.Name,
		Email: profile.Email,
		Phone: profile.PhoneNumber,
		Roles: roles,
	}
}
