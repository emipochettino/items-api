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

func createCategoryHandler(c *gin.Context) *errors.APIError {
	var category *entities.Category
	logger.Info("Creating category")

	if err := c.BindJSON(&category); err != nil {

		logger.Error(fmt.Sprintf("Error binding category request - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info("Category bound")

	result, err := dao.CategoryDAO.CreateCategory(category)
	if err != nil {
		logger.Error("Error binding category request")
		c.JSON(http.StatusBadRequest, err)
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusCreated, result)

	return nil
}

func CreateCategoryHandler(c *gin.Context) {
	errorWrapper(createCategoryHandler, c)
}

func updateCategoryHandler(c *gin.Context) *errors.APIError {
	var category *entities.Category
	logger.Info("Creating category")

	if err := c.BindJSON(&category); err != nil {

		logger.Error(fmt.Sprintf("Error binding category request - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info("Category bound")

	result, err := dao.CategoryDAO.UpdateCategory(category)
	if err != nil {
		logger.Error("Error binding category request")
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func UpdateCategoryHandler(c *gin.Context) {
	errorWrapper(updateCategoryHandler, c)
}

func findCategoryHandler(c *gin.Context) *errors.APIError {
	result, err := dao.CategoryDAO.FindCategories()
	if err != nil {
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func FindCategoryHandler(c *gin.Context) {
	errorWrapper(findCategoryHandler, c)
}

func getCategoryHandler(c *gin.Context) *errors.APIError {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("Error parameter not integer - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info(fmt.Sprintf("Getting category category with id: %d", id))

	result, err := dao.CategoryDAO.GetCategory(int(id))
	if err != nil {
		logger.Error("Error binding category request")
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusCreated, result)
	return nil
}

func GetCategoryHandler(c *gin.Context) {
	errorWrapper(getCategoryHandler, c)
}

func deleteCategoryHandler(c *gin.Context) *errors.APIError {
	logger.Info("Creating category")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(fmt.Sprintf("Error parameter not integer - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	if err := dao.CategoryDAO.DeleteCategory(int(id)); err != nil {
		logger.Error("Error binding category request")
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, nil)
	return nil
}

func DeleteCategoryHandler(c *gin.Context) {
	errorWrapper(deleteCategoryHandler, c)
}
