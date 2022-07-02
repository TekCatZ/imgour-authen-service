package models

type Role int64

const (
	Admin Role = iota
	User
	Disable
)

var (
	Role_name = map[Role]string{
		0: "ADMIN",
		1: "USER",
		2: "DISABLED",
	}
	Role_value = map[string]Role{
		"ADMIN":    Admin,
		"USER":     User,
		"DISABLED": Disable,
	}
)

type UserProfile struct {
	Uid         string `bson:"uid"`
	Name        string `bson:"name"`
	Email       string `bson:"email" validate:"required,email"`
	PhoneNumber string `bson:"phone_number"`
	Roles       []Role `bson:"roles"`
}
