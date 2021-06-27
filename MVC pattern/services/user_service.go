package services

import (
	"../domain"
	"../utils"
)

func GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
