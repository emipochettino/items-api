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

func createItemHandler(c *gin.Context) *errors.APIError {
	var item *entities.Item
	logger.Info("Creating item")

	if err := c.BindJSON(&item); err != nil {

		logger.Error(fmt.Sprintf("Error binding item request - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info("Item bound")

	result, err := dao.ItemDAO.CreateItem(item)
	if err != nil {
		logger.Error("Error binding item request")
		c.JSON(http.StatusBadRequest, err)
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusCreated, result)

	return nil
}

func CreateItemHandler(c *gin.Context) {
	errorWrapper(createItemHandler, c)
}

func updateItemHandler(c *gin.Context) *errors.APIError {
	var item *entities.Item
	logger.Info("Updating item")

	if err := c.BindJSON(&item); err != nil {

		logger.Error(fmt.Sprintf("Error binding item request - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info("Item bound")

	result, err := dao.ItemDAO.UpdateItem(item)
	if err != nil {
		logger.Error("Error binding item request")
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func UpdateItemHandler(c *gin.Context) {
	errorWrapper(updateItemHandler, c)
}

func findItemsHandler(c *gin.Context) *errors.APIError {
	result, err := dao.ItemDAO.FindItems()
	if err != nil {
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func FindItemsHandler(c *gin.Context) {
	errorWrapper(findItemsHandler, c)
}

func getItemHandler(c *gin.Context) *errors.APIError {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("Error parameter not integer - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	logger.Info(fmt.Sprintf("Getting item with id: %d", id))

	result, err := dao.ItemDAO.GetItem(int(id))
	if err != nil {
		logger.Error("Error binding item request")
		return errors.NewBadRequest()
	}

	if result == nil {
		return errors.NewResourceNotFound("Item not found")
	}

	c.JSON(http.StatusCreated, result)
	return nil
}

func GetItemHandler(c *gin.Context) {
	errorWrapper(getItemHandler, c)
}

func deleteItemHandler(c *gin.Context) *errors.APIError {
	logger.Info("Deleting item")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(fmt.Sprintf("Error parameter not integer - error: %s", err.Error()))
		return errors.NewBadRequest()
	}

	if err := dao.ItemDAO.DeleteItem(int(id)); err != nil {
		logger.Error("Error binding item request")
		return errors.NewBadRequest()
	}

	c.JSON(http.StatusOK, nil)
	return nil
}

func DeleteItemHandler(c *gin.Context) {
	errorWrapper(deleteItemHandler, c)
}
