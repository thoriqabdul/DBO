package model

import (
	"fmt"
	"only-test/request"
	"time"
)

type Order struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	CustomerID int       `json:"customer_id"`
	UnitName   string    `json:"unit_name"`
	Price      int       `gorm:"not null" json:"price"`
	Qty        int       `gorm:"not null" json:"qty"`
	Total      int       `gorm:"not null" json:"total"`
	Customer   Customer  `gorm:"references:ID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (mod *Model) CreateOrder(req Order) (Order, error) {
	result := CountTotal(req.Qty, req.Price)
	req.Total = result
	err := mod.db.Create(&req).Error
	if err != nil {
		fmt.Print(err.Error())
		return Order{}, err
	}
	return req, nil
}

func (mod *Model) UpdateOrder(req Order) (Order, error) {
	res, err := mod.DetailOrder(req.ID)
	if err != nil {
		fmt.Print(err.Error())
		return Order{}, err
	}
	res.CustomerID = req.CustomerID
	res.UnitName = req.UnitName
	res.Price = req.Price
	res.Qty = req.Qty
	result := CountTotal(req.Qty, req.Price)
	res.Total = result
	err = mod.db.Save(&res).Error
	if err != nil {
		fmt.Print(err.Error())
		return Order{}, err
	}
	return res, nil
}

func (mod *Model) DetailOrder(id int) (Order, error) {
	var res Order
	err := mod.db.Preload("Customer").First(&res, id).Error
	if err != nil {
		fmt.Print(err.Error())
		return res, err
	}
	return res, nil
}

func (mod *Model) ListOrder(query request.OrderQuery) (res []Order, err error) {
	statment := mod.db.Model(&res)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statment = statment.Joins("JOIN customers c ON c.id = orders.customer_id").Where("unit_name LIKE ?", querySearch).Or("c.name LIKE ?", querySearch).Or("c.email LIKE ?", querySearch)
	}

	err = statment.Limit(query.Limit).
		Offset(query.GetOffset()).
		Preload("Customer").Find(&res).Error

	if err != nil {

		return res, err
	}
	return res, nil
}

func (mod *Model) DeleteOrder(id int) error {
	var res Order
	err := mod.db.Delete(&res, id).Error
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

func CountTotal(qty, price int) int {
	result := qty * price
	return result
}
