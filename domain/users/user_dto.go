package users

import (
	"github.com/vijay-vikram/bookstore_users-api/utils/errors"
	"net/http"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	if user.Email == "" {
		return &errors.RestErr{
			Message: "Email Address is not Valid",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
	}
	return nil
}
