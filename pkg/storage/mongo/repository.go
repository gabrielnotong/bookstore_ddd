package mongo

import (
	"fmt"
	"github.com/gabrielnotong/bookstore/pkg/adding"
	"github.com/gabrielnotong/bookstore/pkg/listing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	*mgo.Database
}

func NewDB(dataSourceName string) (*DB, error) {
	s, err := mgo.Dial(dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = s.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("You are logged in.")

	return &DB{s.DB("bookstore")}, nil
}

func (d *DB) AllBooks() ([]*listing.Book, error) {
	bb := make([]*listing.Book, 0)

	err := d.C("books").Find(bson.M{}).All(&bb)
	if err != nil {
		return nil, err
	}

	return bb, nil
}

func (d *DB) OneBook(isbn string) (*listing.Book, error) {
	b := Book{}
	err := d.C("books").Find(bson.M{"isbn": isbn}).One(&b)
	if err != nil {
		return nil, err
	}

	bo := &listing.Book{
		Isbn:   b.Isbn,
		Title:  b.Title,
		Author: b.Author,
		Price:  b.Price,
	}

	return bo, nil
}

func (d *DB) AddBook(b *adding.Book) error {
	bo := &Book{
		ID:     bson.NewObjectId(),
		Isbn:   b.Isbn,
		Title:  b.Title,
		Author: b.Author,
		Price:  b.Price,
	}

	err := d.C("books").Insert(bo)
	if err != nil {
		return err
	}

	return nil
}