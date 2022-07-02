package db

import (
	"context"
	"errors"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/common"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/models"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserProfile(userProfile models.UserProfile) (*models.UserProfile, error) {
	ctx := context.Background()
	doc, err := userDb.InsertOne(ctx, userProfile)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	uid := doc.InsertedID
	userProfile.Uid = uid.(string)

	return &userProfile, nil
}

func UpdateUserProfile(newUserProfile models.UserProfile) (bool, error) {
	ctx := context.Background()
	uid := newUserProfile.Uid

	updateBsonObject := bson.M{}
	if newUserProfile.Name != "" {
		updateBsonObject["name"] = newUserProfile.Name
	}
	if newUserProfile.Email != "" {
		updateBsonObject["email"] = newUserProfile.Email
	}
	if newUserProfile.PhoneNumber != "" {
		updateBsonObject["phone_number"] = newUserProfile.PhoneNumber
	}
	if len(newUserProfile.Roles) > 0 {
		updateBsonObject["roles"] = newUserProfile.Roles
	}

	err := userDb.UpdateOne(
		ctx,
		bson.M{"uid": uid},
		bson.D{
			{"$set", updateBsonObject},
		},
	)
	if err != nil {
		return false, common.ErrInternal(err)
	}

	return true, nil
}

func FindUserProfileByUid(uid string) (*models.UserProfile, error) {
	ctx := context.Background()
	var err error
	userProfile := models.UserProfile{}
	err = userDb.Find(ctx, bson.M{"uid": uid}).One(&userProfile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrNotFound()
		}
		return nil, common.ErrInternal(err)
	}

	return &userProfile, nil
}

func FindAllUserProfile() ([]models.UserProfile, error) {
	ctx := context.Background()
	var err error
	var userProfiles []models.UserProfile
	err = userDb.Find(ctx, bson.M{}).All(&userProfiles)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrNotFound()
		}
		return nil, common.ErrInternal(err)
	}

	return userProfiles, nil
}

func GetRolesByUid(uid string) ([]models.Role, error) {
	ctx := context.Background()
	var err error
	userProfile := models.UserProfile{}
	err = userDb.Find(ctx, bson.M{"uid": uid}).One(&userProfile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrNotFound()
		}
		return nil, common.ErrInternal(err)
	}

	return userProfile.Roles, nil
}

func AddRolesUserProfile(uid string, newRoles []models.Role) (bool, error) {
	ctx := context.Background()
	emptyErr := errors.New("invalid roles array")
	if len(newRoles) == 0 {
		return false, emptyErr
	}
	err := userDb.UpdateOne(
		ctx,
		bson.M{"uid": uid},
		bson.D{
			{"$set", bson.M{
				"roles": newRoles,
			}},
		},
	)
	if err != nil {
		return false, common.ErrInternal(err)
	}

	return true, nil
}

func RemoveRolesUserProfile(uid string) (bool, error) {
	ctx := context.Background()
	var err error
	err = userDb.UpdateOne(
		ctx,
		bson.M{"uid": uid},
		bson.D{
			{"$set", bson.M{
				"roles": []models.Role{},
			}},
		},
	)
	if err != nil {
		return false, common.ErrInternal(err)
	}

	return true, nil
}
