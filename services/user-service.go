package services

import (
	"github.com/bhawana/bookstore_user-apis/domain/users"
	"github.com/bhawana/bookstore_user-apis/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
