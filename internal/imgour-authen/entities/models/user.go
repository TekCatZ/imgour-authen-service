package models

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
