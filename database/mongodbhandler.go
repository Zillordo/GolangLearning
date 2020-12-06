package database

import (
	model "../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongodbHandler(connection string) (*MongodbHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}

func (h *MongodbHandler) GetDynos() ([]model.Animal, error) {
	s := h.getFreshSession()
	defer s.Close()
	var animal []model.Animal
	err := h.db(s).Find(nil).All(&animal)
	return animal, err
}

func (h *MongodbHandler) GetDynoByNickname(nickname string) (model.Animal, error) {
	s := h.getFreshSession()
	defer s.Close()
	var a model.Animal
	err := h.db(s).Find(bson.M{"nickname": nickname}).One(&a)
	return a, err
}

func (h *MongodbHandler) GetDynosByType(dinoType string) ([]model.Animal, error) {
	s := h.getFreshSession()
	defer s.Close()
	var animals []model.Animal
	err := h.db(s).Find(bson.M{"animal_type": dinoType}).All(&animals)
	return animals, err
}

func (h *MongodbHandler) AddAnimal(a model.Animal) error {
	s := h.getFreshSession()
	defer s.Close()
	return h.db(s).Insert(a)
}

func (h *MongodbHandler) UpdateAnimal(a model.Animal, nickname string) error {
	s := h.getFreshSession()
	defer s.Close()
	return h.db(s).Update(bson.M{"nickname": nickname}, a)
}

func (h *MongodbHandler) getFreshSession() *mgo.Session {
	return h.Session.Copy()
}

func (h *MongodbHandler) db(session *mgo.Session) *mgo.Collection {
	return session.DB("Dino").C("animals")
}
