package database

import (
	model "../models"
	"errors"
)

type DbType uint8

const (
	MYSQL DbType = iota
	POSTGRES
	MONGODB
)

type DinoDBHandler interface {
	GetDynos() ([]model.Animal, error)
	GetDynoByNickname(nickname string) (model.Animal, error)
	GetDynosByType(dinoType string) ([]model.Animal, error)
	AddAnimal(a model.Animal) error
	UpdateAnimal(a model.Animal, nickname string) error
}

var DBTypeNotSupported = errors.New("the Database type provided is not supported")

func GetDatabaseHandler(dbType DbType, connection string) (DinoDBHandler, error) {
	switch dbType {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongodbHandler(connection)
	case POSTGRES:
		return NewPQHandler(connection)
	}
	return nil, DBTypeNotSupported
}
