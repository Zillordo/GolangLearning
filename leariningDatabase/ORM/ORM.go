package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type animal struct {
	gorm.Model
	AnimalType string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zone       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

func main() {
	db, err := gorm.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{})
	a := animal{
		AnimalType: "Some type",
		Nickname:   "nick",
		Zone:       1,
		Age:        300,
	}
	db.Save(&a)
}
