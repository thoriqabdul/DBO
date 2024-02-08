package view

import "only-test/model"

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
