package database

import (
	model "../models"
	"database/sql"
	"fmt"
	"log"
)

type SQLHandler struct {
	*sql.DB
}

func (handler *SQLHandler) GetDynos() ([]model.Animal, error) {
	return handler.sendQuery("SELECT * from animals")
}

func (handler *SQLHandler) GetDynoByNickname(nickname string) (model.Animal, error) {
	row := handler.QueryRow(fmt.Sprintf("SELECT * from animals where nickname = '%s'", nickname))

	var a model.Animal
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (handler *SQLHandler) GetDynosByType(dinoType string) ([]model.Animal, error) {
	return handler.sendQuery(fmt.Sprintf("select * from Animals where Animal_type = '%s'", dinoType))
}

func (handler *SQLHandler) AddAnimal(a model.Animal) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into Animals (Animal_type,nickname,zone,age) values ('%s','%s',%d,%d",
		a.AnimalType, a.Nickname, a.Zone, a.Age))
	return err
}

func (handler *SQLHandler) UpdateAnimal(a model.Animal, nickname string) error {
	_, err := handler.Exec(fmt.Sprintf(
		"Update Animals set Animal_type = '%s', nickname = '%s', zone = %d, age = %d where nickname = '%s' and ID = '%s'",
		a.AnimalType, a.Nickname, a.Zone, a.Age, nickname, a.ID))
	return err
}

func (handler *SQLHandler) sendQuery(q string) ([]model.Animal, error) {
	var Animals []model.Animal
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a model.Animal
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}
