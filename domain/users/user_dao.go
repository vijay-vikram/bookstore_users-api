package users

import (
	"fmt"
	"github.com/vijay-vikram/bookstore_users-api/utils/dates"
	"github.com/vijay-vikram/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	results := usersDB[user.Id]
	if results == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found ", user.Id))
	}
	user.Id = results.Id
	user.Email = results.Email
	user.LastName = results.LastName
	user.FirstName = results.FirstName
	user.DateCreated = results.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exits", user.Id))
	}
	user.DateCreated = dates.GetNoWString()
	usersDB[user.Id] = user
	return nil
}
