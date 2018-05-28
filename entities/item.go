package entities

import (
	"time"
)

var UserLocalMysql = "root"
var PassLocalMysql = ""
var DataBaseLocalMysql = "test"

/*
type Item struct {
	ID              int
	Name            string
	Price           float64
	Categories      []Category
	User            User
	DateCreated     *time.ISO8601Time
	DateLastUpdated *time.ISO8601Time
}

type User struct {
	ID              int
	Username        string
	Password        string
	Email           string
	phoneNumber     string
	DateCreated     *time.ISO8601Time
	DateLastUpdated *time.ISO8601Time
}
*/
type Category struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"varchar(255)" json:"name"`
	CreatedAt time.Time `json:"date_created"`
	UpdatedAt time.Time `json:"date_last_updated"`
}
