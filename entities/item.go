package entities

import (
	"time"
)

var UserLocalMysql = "root"
var PassLocalMysql = ""
var DataBaseLocalMysql = "test"

type Item struct {
	ID         int        `gorm:"primary_key" json:"id"`
	Name       string     `gorm:"varchar(255)" json:"name"`
	Price      float64    `json:"price"`
	Categories []Category `gorm:"many2many:item_category;" json:"categories,omitempty"`
	User       *User      `json:"user,omitempty"`
	UserId     int        `json:"-"`
	CreatedAt  time.Time  `json:"date_created"`
	UpdatedAt  time.Time  `json:"date_last_updated"`
}

type User struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Username    string    `gorm:"varchar(255)" json:"username"`
	Password    string    `gorm:"varchar(255)" json:"password"`
	Email       string    `gorm:"varchar(255)" json:"email"`
	PhoneNumber string    `gorm:"varchar(15)" json:"phone_number"`
	CreatedAt   time.Time `json:"date_created"`
	UpdatedAt   time.Time `json:"date_last_updated"`
}

type Category struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"varchar(255)" json:"name"`
	CreatedAt time.Time `json:"date_created"`
	UpdatedAt time.Time `json:"date_last_updated"`
}
