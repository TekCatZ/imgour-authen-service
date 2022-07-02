package profile

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/common"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/models/dto"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"net/http"
)

func GetProfileHandler(c *gin.Context) {
	sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())

	userID := sessionContainer.GetUserID()

	user, err := repositories.GetUser(userID)
	if err != nil {
		log.Error("Error getting user: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": common.ErrInternal(nil),
		})
		return
	}

	profile := dto.ParseFromModel(user)
	c.JSON(http.StatusOK, profile.ParseToResponse())
}
