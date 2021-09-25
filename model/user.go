package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func (u User) TableName() string {
	return "user"
}

func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}

func (u *User) Create() error {
	result := GetDb().Create(u)
	return result.Error
}
