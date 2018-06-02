package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emipochettino/items-api-go/dao"
	"github.com/emipochettino/items-api-go/entities"
	"github.com/emipochettino/items-api-go/errors"
	"github.com/emipochettino/items-api-go/logger"
	"github.com/gin-gonic/gin"
)

func createUserHandler(c *gin.Context) *errors.APIError {
	var user *entities.User
	logger.Info("Creating user")

	if err := c.BindJSON(&user); err != nil {

		logger.Error(fmt.Sprintf("Error binding user request - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info("User bound")

	result, err := dao.UserDAO.CreateUser(user)
	if err != nil {
		logger.Error("Error binding user request")
		c.JSON(http.StatusBadRequest, err)
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusCreated, result)

	return nil
}

func CreateUserHandler(c *gin.Context) {
	errorWrapper(createUserHandler, c)
}

func updateUserHandler(c *gin.Context) *errors.APIError {
	var user *entities.User
	logger.Info("Updating user")

	if err := c.BindJSON(&user); err != nil {

		logger.Error(fmt.Sprintf("Error binding user request - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info("User bound")

	result, err := dao.UserDAO.UpdateUser(user)
	if err != nil {
		logger.Error("Error binding user request")
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func UpdateUserHandler(c *gin.Context) {
	errorWrapper(updateUserHandler, c)
}

func findUsersHandler(c *gin.Context) *errors.APIError {
	result, err := dao.UserDAO.FindUsers()
	if err != nil {
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func FindUsersHandler(c *gin.Context) {
	errorWrapper(findUsersHandler, c)
}

func getUserHandler(c *gin.Context) *errors.APIError {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("Error parameter not integer - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info(fmt.Sprintf("Getting user with id: %d", id))

	result, err := dao.UserDAO.GetUser(int(id))
	if err != nil {
		logger.Error("Error binding user request")
		return errors.NewBadRequest()
	}

	if result == nil {
		return errors.NewResourceNotFound("User not found")
	}

	c.JSON(http.StatusCreated, result)
	return nil
}

func GetUserHandler(c *gin.Context) {
	errorWrapper(getUserHandler, c)
}

func deleteUserHandler(c *gin.Context) *errors.APIError {
	logger.Info("Deleting user")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(fmt.Sprintf("Error parameter not integer - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	if err := dao.UserDAO.DeleteUser(int(id)); err != nil {
		logger.Error("Error binding user request")
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, nil)
	return nil
}

func DeleteUserHandler(c *gin.Context) {
	errorWrapper(deleteUserHandler, c)
}
