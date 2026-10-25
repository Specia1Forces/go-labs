package main

import "fmt"

type Book struct {
	ID        int
	Title     string
	Author    string
	Year      int
	Available bool
}

func NewBook(id int, title string, author string, year int, available bool) Book {
	return Book{
		ID:        id,
		Title:     title,
		Author:    author,
		Year:      year,
		Available: available,
	}
}

// Description uses a value receiver because it only reads the book.
func (book Book) Description() string {
	status := "borrowed"
	if book.Available {
		status = "available"
	}

	return fmt.Sprintf("%d: %s by %s (%d) — %s", book.ID, book.Title, book.Author, book.Year, status)
}

// Borrow uses a pointer receiver because it changes the original book.
func (book *Book) Borrow() bool {
	if !book.Available {
		return false
	}

	book.Available = false
	return true
}

func FindBooksByAuthor(books []Book, author string) []Book {
	foundBooks := make([]Book, 0)

	for _, book := range books {
		if book.Author == author {
			foundBooks = append(foundBooks, book)
		}
	}

	return foundBooks
}

func CountAvailableBooks(books []Book) int {
	count := 0

	for _, book := range books {
		if book.Available {
			count++
		}
	}

	return count
}

func FindNewestBook(books []Book) (Book, bool) {
	if len(books) == 0 {
		return Book{}, false
	}

	newestBook := books[0]
	for _, book := range books[1:] {
		if book.Year > newestBook.Year {
			newestBook = book
		}
	}

	return newestBook, true
}

func GroupBooksByAuthor(books []Book) map[string][]Book {
	groups := make(map[string][]Book)

	for _, book := range books {
		groups[book.Author] = append(groups[book.Author], book)
	}

	return groups
}
