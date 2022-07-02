package repositories

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/models"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/db"
)

func SaveUser(user *models.UserProfile) (string, error) {
	// Save user to database
	profile, err := db.CreateUserProfile(*user)
	if err != nil {
		return "", err
	}

	return profile.Uid, nil
}

func GetUser(uid string) (*models.UserProfile, error) {
	// Get user from database
	profile, err := db.FindUserProfileByUid(uid)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func GetRolesByUid(uid string) ([]models.Role, error) {
	// Get user from database
	roles, err := db.GetRolesByUid(uid)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
