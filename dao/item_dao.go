package dao

import (
	"fmt"

	"github.com/emipochettino/items-api-go/entities"
	"github.com/emipochettino/items-api-go/logger"
)

type IItemDao interface {
	CreateItem(item *entities.Item) (*entities.Item, error)
	UpdateItem(item *entities.Item) (*entities.Item, error)
	DeleteItem(id int) error
	FindItems() (*[]entities.Item, error)
	GetItem(id int) (*entities.Item, error)
}

type ItemDao struct{}

func (ItemDao) CreateItem(item *entities.Item) (*entities.Item, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Create(item).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to save a new item - error: %s", err.Error()))
		return nil, err
	}

	return item, nil
}

func (ItemDao) UpdateItem(item *entities.Item) (*entities.Item, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Model(&entities.Item{}).Update(item).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to update a item - error: %s", err.Error()))
		return nil, err
	}

	return item, nil
}

func (ItemDao) DeleteItem(id int) error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	item, err := ItemDAO.GetItem(id)
	if err != nil {
		return err
	}

	if err := db.Delete(item).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to save a new item - error: %s", err.Error()))
		return err
	}

	return nil
}

func (ItemDao) FindItems() (*[]entities.Item, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	items := &[]entities.Item{}
	if err := db.Find(items).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when finding users - error: %s", err.Error()))
		return nil, err
	}

	return items, nil
}

func (ItemDao) GetItem(id int) (*entities.Item, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	item := &entities.Item{}
	if db.Model(item).Preload("Categories").Preload("User").Where(id).Find(item).RecordNotFound() {
		return nil, nil
	}

	return item, nil
}
