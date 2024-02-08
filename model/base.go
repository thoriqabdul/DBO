package model

import (
	"only-test/utils"
)

type Model struct {
	db *utils.Database
}

func NewModel(db *utils.Database) Model {

	return Model{db: db}
}
