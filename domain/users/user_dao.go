package users

import (
	"fmt"
	"github.com/vijay-vikram/bookstore_users-api/datasources/mysql/users_db"
	"github.com/vijay-vikram/bookstore_users-api/logger"
	"github.com/vijay-vikram/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUser   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when try to prepare get user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when try to get user by ID", getErr)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when try to prepare Save user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when try to save user", saveErr)
		return errors.NewInternalServerError("Database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when try to get last insert ID after creating a new user", err)
		return errors.NewInternalServerError("Database error")
	}

	user.Id = userID

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when try to prepare Update user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when try to Update user", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when try to prepare Delete user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when try to Delete user", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUser)
	if err != nil {
		logger.Error("error when try to prepare Find User By Status statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when try to Find user by status", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer rows.Close()

	var userList = make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			logger.Error("error when try to scan user row in user struct", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		userList = append(userList, user)
	}

	if len(userList) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching given status %s", status))
	}
	return userList, nil
}
