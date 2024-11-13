package handlers

import (
	"github.com/GnauqTheBeast/models"
	"html/template"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// List all books
func ListBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	db.Find(&books)
	tmpl.ExecuteTemplate(w, "index.html", books)
}

// Render form to create a new book
func NewBook(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "new.html", nil)
}

// Create a new book
func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		author := r.FormValue("author")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)

		db.Create(&models.Book{Title: title, Author: author, Price: price})
		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

// Edit a book
func EditBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var book models.Book
	db.First(&book, id)
	tmpl.ExecuteTemplate(w, "edit.html", book)
}

// Update a book
func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		title := r.FormValue("title")
		author := r.FormValue("author")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)

		var book models.Book
		db.First(&book, id)
		book.Title = title
		book.Author = author
		book.Price = price
		db.Save(&book)

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

// Delete a book
func DeleteBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Delete(&models.Book{}, id)
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
