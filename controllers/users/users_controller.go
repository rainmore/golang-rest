package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rainmore.com.au/rest-api/domain/errors"
	"rainmore.com.au/rest-api/domain/users"
	"rainmore.com.au/rest-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestError{
			Message: "Invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}

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

func GetUser(c *gin.Context) {
	userId, userError := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userError != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO")
}

func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO")
}
