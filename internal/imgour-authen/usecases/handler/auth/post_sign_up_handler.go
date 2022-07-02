package auth

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/models"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/repositories"
	"github.com/sirupsen/logrus"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartypasswordless/tplmodels"
)

type PostSignUpHandler interface {
	Handle(user tplmodels.User)
}

type signUpHandler struct{}

func (*signUpHandler) Handle(user tplmodels.User) {
	userDao := &models.UserProfile{
		Uid:         user.ID,
		Name:        "",
		Email:       *user.Email,
		PhoneNumber: *user.PhoneNumber,
		Roles: []models.Role{
			models.User,
		},
	}
	_, err := repositories.SaveUser(userDao)
	if err != nil {
		logrus.Errorln("Error saving user: ", err)
		return
	}
}

func GetPostSignUpHandler() *signUpHandler {
	return &signUpHandler{}
}
