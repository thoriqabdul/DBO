package model

import (
	"fmt"
	"only-test/request"
	"time"
)

type Customer struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `gorm:"not null;unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (mod *Model) CreateCustomer(req Customer) (Customer, error) {
	err := mod.db.Create(&req).Error
	if err != nil {
		fmt.Print(err.Error())
		return Customer{}, err
	}
	return req, nil
}

func (mod *Model) UpdateCustomer(req Customer) (Customer, error) {
	res, err := mod.DetailCustomer(req.ID)
	if err != nil {
		fmt.Print(err.Error())
		return Customer{}, err
	}
	res.Email = req.Email
	res.Name = req.Name
	err = mod.db.Save(&res).Error
	if err != nil {
		fmt.Print(err.Error())
		return Customer{}, err
	}
	return res, nil
}

func (mod *Model) DetailCustomer(id int) (Customer, error) {
	var res Customer
	err := mod.db.First(&res, id).Error
	if err != nil {
		fmt.Print(err.Error())
		return res, err
	}
	return res, nil
}

func (mod *Model) DeleteCustomer(id int) error {
	var res Customer
	err := mod.db.Delete(&res, id).Error
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

func (mod *Model) ListCustomer(query request.CustomerQuery) (res []Customer, err error) {
	statment := mod.db.Model(&res)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statment = statment.Where("name LIKE ?", querySearch).Or("email LIKE ?", querySearch)
	}

	err = statment.Limit(query.Limit).
		Offset(query.GetOffset()).
		Find(&res).Error

	if err != nil {

		return res, err
	}
	return res, nil
}
