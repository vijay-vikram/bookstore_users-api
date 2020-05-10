package users

import (
	"github.com/gin-gonic/gin"
	"github.com/vijay-vikram/bookstore_users-api/domain/users"
	"github.com/vijay-vikram/bookstore_users-api/services"
	"github.com/vijay-vikram/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, userIdError := getUserId(c.Param("user_id"))
	if userIdError != nil {
		c.JSON(userIdError.Status, userIdError)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, userIdError := getUserId(c.Param("user_id"))
	if userIdError != nil {
		c.JSON(userIdError.Status, userIdError)
		return
	}

	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, restErr := services.UpdateUser(user, isPartial)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, userIdError := getUserId(c.Param("user_id"))
	if userIdError != nil {
		c.JSON(userIdError.Status, userIdError)
		return
	}

	restErr := services.DeleteUser(userId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	usersList, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, usersList)
}

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userIdError := strconv.ParseInt(userIdParam, 10, 64)
	if userIdError != nil {
		return 0, errors.NewBadRequestError("User Id should be a Integer")
	}
	return userId, nil
}
