package model

import (
	"fmt"
	"only-test/request"
	"only-test/utils"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (mod *Model) CreateUser(req User) (User, error) {
	hasPassword, err := utils.GenerateHashFromString(req.Password)
	if err != nil {
		return User{}, err
	}
	req.Password = hasPassword
	err = mod.db.Create(&req).Error
	if err != nil {
		return User{}, err
	}
	return req, nil
}

func (mod *Model) CheckUser(req request.LoginRequest) (User, error, bool) {
	var res User
	err := mod.db.Where("email=?", req.Email).First(&res).Error
	if err != nil {
		fmt.Print(err.Error())
		return res, err, false
	}
	return res, nil, true
}
