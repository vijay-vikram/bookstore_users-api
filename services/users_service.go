package services

import (
	"github.com/vijay-vikram/bookstore_users-api/domain/users"
	"github.com/vijay-vikram/bookstore_users-api/utils/crypto_utils"
	"github.com/vijay-vikram/bookstore_users-api/utils/dates"
	"github.com/vijay-vikram/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := users.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = dates.GetNoWDbFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.Email != "" {
			currentUser.Email = user.Email
		}
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
	} else {
		currentUser.Email = user.Email
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
	}

	restErr := currentUser.Update()
	if restErr != nil {
		return nil, restErr
	}

	return currentUser, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
