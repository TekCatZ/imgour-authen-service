package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type Role int64

const (
	Admin Role = iota
	User
)

type UserProfile struct {
	Uid         string `bson:"uid"`
	Name        string `bson:"name"`
	Email       string `bson:"email" validate:"required,email"`
	PhoneNumber string `bson:"phone_number"`
	Roles       []Role `bson:"roles"`
}

func CreateUserProfile(userProfile UserProfile) (*UserProfile, error) {
	ctx := context.Background()
	doc, err := userDb.InsertOne(ctx, userProfile)
	if err != nil {
		return nil, err
	}

	uid := doc.InsertedID
	userProfile.Uid = uid.(string)

	return &userProfile, nil
}

func UpdateUserProfile(newUserProfile UserProfile) (bool, error) {
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
		return false, err
	}

	return true, nil
}

func FindUserProfileByUid(uid string) (*UserProfile, error) {
	ctx := context.Background()
	var err error
	userProfile := UserProfile{}
	err = userDb.Find(ctx, bson.M{"uid": uid}).One(&userProfile)
	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}

func FindAllUserProfile() ([]UserProfile, error) {
	ctx := context.Background()
	var err error
	userProfiles := []UserProfile{}
	err = userDb.Find(ctx, bson.M{}).All(&userProfiles)
	if err != nil {
		return nil, err
	}

	return userProfiles, nil
}

func GetRolesByUid(uid string) ([]Role, error) {
	ctx := context.Background()
	var err error
	userProfile := UserProfile{}
	err = userDb.Find(ctx, bson.M{"uid": uid}).One(&userProfile)
	if err != nil {
		return nil, err
	}

	return userProfile.Roles, nil
}

func UpdateRolesUserProfile(uid string, newRoles []Role) (bool, error) {
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
		return false, err
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
				"roles": []Role{},
			}},
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}
