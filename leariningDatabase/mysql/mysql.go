package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	//connect to the leariningDatabase
	db, err := sql.Open("mysql", "root:root@/Dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//general query with arguments
	rows, err := db.Query("select * from Dino.animals where age > ?", 10)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var animals []animal
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(animals)

	//query a single row
	row := db.QueryRow("select * from Dino.animals where age > 10")
	var a animal
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}

	//insert row
	res, err := db.Exec("INSERT into Dino.animals (animal_type,nickname,zone,age) values ('Centaurus', 'Caron', 3, 22)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	//updating rows
	res, err = db.Exec("UPDATE Dino.animals set age = ? where id = ?", 14, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}
