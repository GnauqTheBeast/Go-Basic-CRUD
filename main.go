package main

import (
	"fmt"
	"github.com/GnauqTheBeast/handlers"
	"github.com/GnauqTheBeast/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		return
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.ListBooks(db, w, r)
		})
		r.Get("/new", handlers.NewBook)
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
			handlers.CreateBook(db, w, r)
		})
		r.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
			handlers.EditBook(db, w, r)
		})
		r.Post("/update", func(w http.ResponseWriter, r *http.Request) {
			handlers.UpdateBook(db, w, r)
		})
		r.Get("/delete", func(w http.ResponseWriter, r *http.Request) {
			handlers.DeleteBook(db, w, r)
		})
	})

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
