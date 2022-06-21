package repositories

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/models"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/db"
)

func SaveUser(user models.UserProfile) (string, error) {
	// Save user to database
	profile, err := db.CreateUserProfile(user)
	if err != nil {
		return "", err
	}

	return profile.Uid, nil
}
