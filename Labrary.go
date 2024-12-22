package main

import (
	"fmt"
)

func main() {
	library := Library{Books: make(map[string]Book)}
	var id, title, author string
	fmt.Print("Enter Book ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Book Title: ")
	fmt.Scan(&title)
	fmt.Print("Enter Book Author: ")
	fmt.Scan(&author)
	library.AddBook(Book{ID: id, Title: title, Author: author})
}

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
	fmt.Println("Book added successfully.")
}
