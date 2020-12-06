package dynomodel

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"Animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}
