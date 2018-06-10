package dao

import (
	"fmt"

	"github.com/emipochettino/items-api-go/entities"
	"github.com/emipochettino/items-api-go/logger"
)

type IUserDao interface {
	CreateUser(user *entities.User) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(id int) error
	FindUsers() (*[]entities.User, error)
	GetUsers(id int) (*entities.User, error)
}

type UserDao struct{}

func (UserDao) CreateUser(user *entities.User) (*entities.User, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Create(user).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to save a new user - error: %s", err.Error()))
		return nil, err
	}

	return user, nil
}

func (UserDao) UpdateUser(user *entities.User) (*entities.User, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Model(&entities.User{}).Update(user).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to update a user - error: %s", err.Error()))
		return nil, err
	}

	return user, nil
}

func (UserDao) DeleteUser(id int) error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	user, err := UserDAO.GetUser(id)
	if err != nil {
		return err
	}

	if err := db.Delete(user).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when it tried to save a new user - error: %s", err.Error()))
		return err
	}

	return nil
}

func (UserDao) FindUsers() (*[]entities.User, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	users := &[]entities.User{}
	if err := db.Find(users).Error; err != nil {
		logger.Error(fmt.Sprintf("Something went wrong when finding users - error: %s", err.Error()))
		return nil, err
	}

	return users, nil
}

func (UserDao) GetUser(id int) (*entities.User, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := &entities.User{}
	if db.Where(id).First(user).RecordNotFound() {
		return nil, nil
	}

	return user, nil
}
