package adding

type Book struct {
	Isbn   string        `json:"isbn" bson:"isbn"`
	Title  string        `json:"title" bson:"title"`
	Author string        `json:"author" bson:"author"`
	Price  float32       `json:"price" bson:"price"`
}
