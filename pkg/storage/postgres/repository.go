package postgres

import (
	"database/sql"
	"fmt"
	"github.com/gabrielnotong/bookstore/pkg/adding"
	"github.com/gabrielnotong/bookstore/pkg/listing"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("You are logged in.")

	return &DB{db}, nil
}

func (d *DB) AllBooks() ([]*listing.Book, error) {
	rows, err := d.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bb := make([]*listing.Book, 0)

	for rows.Next() {
		bo := &Book{}
		err := rows.Scan(&bo.Isbn, &bo.Title, &bo.Author, &bo.Price)
		if err != nil {
			return nil, err
		}

		b := &listing.Book{
			Isbn:   bo.Isbn,
			Title:  bo.Title,
			Author: bo.Author,
			Price:  bo.Price,
		}

		bb = append(bb, b)
	}

	return bb, nil
}

func (d *DB) OneBook(isbn string) (*listing.Book, error) {
	row := d.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	b := Book{}
	err := row.Scan(&b.ID, &b.Isbn, &b.Title, &b.Author, &b.Price)
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
	_, err := d.Exec(
		"INSERT INTO books(isbn, title, author, price) VALUES ($1, $2, $3, $4)",
		b.Isbn, b.Title, b.Author, b.Price)
	if err != nil {
		return err
	}

	return nil
}
