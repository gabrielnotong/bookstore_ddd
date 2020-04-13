package server

import (
	"encoding/json"
	"fmt"
	"github.com/gabrielnotong/bookstore/pkg/adding"
	"github.com/gabrielnotong/bookstore/pkg/listing"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Handler(ls listing.Service, as adding.Service) *httprouter.Router {
	router := httprouter.New()

	router.GET("/books", AllBooks(ls))
	router.GET("/books/book/:isbn", OneBook(ls))
	router.POST("/books/add", AddBook(as))

	return router
}

func AllBooks(s listing.Service) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		bb, err := s.AllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, b := range bb {
			_, _ = fmt.Fprintf(w, "- %s, %s, %s, $%.2f\n", b.Isbn, b.Title, b.Author, b.Price)
		}

		//w.Header().Set("Content-Type", "application/json")
		//_ = json.NewEncoder(w).Encode(bb)
	}
}

func OneBook(s listing.Service) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		isbn := p.ByName("isbn")
		if isbn == "" {
			http.Error(w, "ISBN is mandatory", http.StatusInternalServerError)
			return
		}

		b, err := s.OneBook(isbn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = fmt.Fprintf(w, "- %s, %s, %s, $%.2f\n", b.Isbn, b.Title, b.Author, b.Price)

		//w.Header().Set("Content-Type", "application/json")
		//_ = json.NewEncoder(w).Encode(b)
	}
}

func AddBook(s adding.Service) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		b := &adding.Book{}

		err := json.NewDecoder(req.Body).Decode(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.AddBook(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = fmt.Fprintf(w, "- Added: %s, %s, %s, $%.2f\n", b.Isbn, b.Title, b.Author, b.Price)

		//w.Header().Set("Content-Type", "application/json")
		//_ = json.NewEncoder(w).Encode(b)
	}
}
