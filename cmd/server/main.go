package main

import (
	"github.com/gabrielnotong/bookstore/pkg/adding"
	"github.com/gabrielnotong/bookstore/pkg/http/server"
	"github.com/gabrielnotong/bookstore/pkg/listing"
	"github.com/gabrielnotong/bookstore/pkg/storage/postgres"
	"net/http"
)

func main() {
	db, err := postgres.NewDB("postgres://postgres:password@localhost/bookstore?sslmode=disable")
	//db, err := mongo.NewDB("mongodb://mongo:mongo@localhost/bookstore")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ls := listing.NewService(db)
	as := adding.NewService(db)
	r := server.Handler(ls, as)
	_ = http.ListenAndServe(":8080", r)
}
