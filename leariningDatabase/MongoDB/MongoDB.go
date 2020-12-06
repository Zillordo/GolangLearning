package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type animal struct {
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	animalCollection := session.DB("Dino").C("animals")
	err = animalCollection.DropCollection()
	if err != nil {
		log.Fatal(err)
	}

	animals := []interface{}{
		animal{
			AnimalType: "first",
			Nickname:   "first",
			Zone:       1,
			Age:        87,
		},
		animal{
			AnimalType: "second",
			Nickname:   "second",
			Zone:       4,
			Age:        33,
		},
		animal{
			AnimalType: "third",
			Nickname:   "third",
			Zone:       2,
			Age:        421,
		},
	}

	//inserting slice of values into collection
	err = animalCollection.Insert(animals...)
	if err != nil {
		log.Fatal(err)
	}

	//updating row in collection where nickname = third
	err = animalCollection.Update(bson.M{"nickname": "third"}, bson.M{"$set": bson.M{"age": 18}})
	if err != nil {
		log.Fatal(err)
	}

	//removing second row
	err = animalCollection.Remove(bson.M{"nickname": "second"})
	if err != nil {
		log.Fatal(err)
	}

	//query to data with conditions
	query := bson.M{
		"age": bson.M{
			"$gt": 10,
		},
		"zone": bson.M{
			"$in": []int{1, 2},
		},
	}
	var res []animal
	err = animalCollection.Find(query).All(&res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
