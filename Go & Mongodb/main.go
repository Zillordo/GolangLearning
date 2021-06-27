package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"./controllers"
)

func main() {
	router := httprouter.New()
	userController := controllers.NewUserController(getDatabase())
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user/:id", userController.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDatabase() *mgo.Database {
	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}
	return s.DB("wdg")
}
