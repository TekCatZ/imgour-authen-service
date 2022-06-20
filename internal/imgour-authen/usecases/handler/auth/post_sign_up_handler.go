package auth

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/repositories"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/db"
	"github.com/sirupsen/logrus"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartypasswordless/tplmodels"
)

type PostSignUpHandler interface {
	Handle(user tplmodels.User)
}

type signUpHandler struct{}

func (*signUpHandler) Handle(user tplmodels.User) {
	userDao := &db.UserProfile{
		Uid:         user.ID,
		Name:        "",
		Email:       *user.Email,
		PhoneNumber: *user.PhoneNumber,
		Roles: []db.Role{
			db.User,
		},
	}
	_, err := repositories.SaveUser(*userDao)
	if err != nil {
		logrus.Errorln("Error saving user: ", err)
		return
	}
}

func GetPostSignUpHandler() *signUpHandler {
	return &signUpHandler{}
}
