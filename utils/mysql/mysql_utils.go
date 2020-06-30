package mysql

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/vijay-vikram/bookstore_users-api/utils/errors"
	"strings"
)

const (
	ErrorNoRow = "no rows in result set"
)

func ParseError(error error) *errors.RestErr {

	sqlError, ok := error.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(error.Error(), ErrorNoRow) {
			return errors.NewNotFoundError(fmt.Sprintf("no record mtching given id"))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error parsing database response"))
	}

	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("Invalid data"))
	}

	return errors.NewInternalServerError(fmt.Sprintf("error processing request"))
}
