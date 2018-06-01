package dao

import (
	"fmt"

	"github.com/emipochettino/items-api-go/entities"
	"github.com/emipochettino/items-api-go/logger"
)

type ICategoryDao interface {
	CreateCategory(category *entities.Category) (*entities.Category, error)
	UpdateCategory(category *entities.Category) (*entities.Category, error)
	DeleteCategory(id int) error
	FindCategories() (*[]entities.Category, error)
	GetCategory(id int) (*entities.Category, error)
}

type CategoryDao struct{}

func (CategoryDao) CreateCategory(category *entities.Category) (*entities.Category, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Create(category).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to save a new category - error: %s", err.Error()))
		return nil, err
	}

	return category, nil
}

func (CategoryDao) UpdateCategory(category *entities.Category) (*entities.Category, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Model(&entities.Category{}).Update(category).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to update a category - error: %s", err.Error()))
		return nil, err
	}

	return category, nil
}

func (CategoryDao) DeleteCategory(id int) error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	category, err := CategoryDAO.GetCategory(id)
	if err != nil {
		return err
	}

	if err := db.Delete(category).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to save a new category - error: %s", err.Error()))
		return err
	}

	return nil
}

func (CategoryDao) FindCategories() (*[]entities.Category, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	categories := &[]entities.Category{}
	if err := db.Find(categories).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when finding categories - error: %s", err.Error()))
		return nil, err
	}

	return categories, nil
}

func (CategoryDao) GetCategory(id int) (*entities.Category, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	category := &entities.Category{}
	if db.Where(&entities.Category{ID: id}).First(category).RecordNotFound() {
		return nil, nil
	}

	return category, nil
}
