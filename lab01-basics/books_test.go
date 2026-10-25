package main

import (
	"reflect"
	"testing"
)

func TestFindBooksByAuthor(t *testing.T) {
	books := []Book{
		NewBook(1, "1984", "George Orwell", 1949, true),
		NewBook(2, "Animal Farm", "George Orwell", 1945, false),
		NewBook(3, "Effective Java", "Joshua Bloch", 2018, true),
	}

	tests := []struct {
		name   string
		author string
		want   []Book
	}{
		{
			name:   "author has two books",
			author: "George Orwell",
			want:   []Book{books[0], books[1]},
		},
		{
			name:   "author has one book",
			author: "Joshua Bloch",
			want:   []Book{books[2]},
		},
		{
			name:   "author is not found",
			author: "Unknown",
			want:   []Book{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FindBooksByAuthor(books, test.author)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("FindBooksByAuthor() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestCountAvailableBooks(t *testing.T) {
	tests := []struct {
		name  string
		books []Book
		want  int
	}{
		{name: "empty catalog", books: []Book{}, want: 0},
		{
			name: "some books are available",
			books: []Book{
				NewBook(1, "First", "Author", 2020, true),
				NewBook(2, "Second", "Author", 2021, false),
				NewBook(3, "Third", "Author", 2022, true),
			},
			want: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CountAvailableBooks(test.books)
			if got != test.want {
				t.Fatalf("CountAvailableBooks() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestFindNewestBook(t *testing.T) {
	tests := []struct {
		name      string
		books     []Book
		want      Book
		wantFound bool
	}{
		{
			name:      "newest book is found",
			books:     []Book{NewBook(1, "Old", "Author", 1990, true), NewBook(2, "New", "Author", 2024, true)},
			want:      NewBook(2, "New", "Author", 2024, true),
			wantFound: true,
		},
		{
			name:      "empty catalog",
			books:     []Book{},
			want:      Book{},
			wantFound: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, found := FindNewestBook(test.books)
			if found != test.wantFound || got != test.want {
				t.Fatalf("FindNewestBook() = (%v, %t), want (%v, %t)", got, found, test.want, test.wantFound)
			}
		})
	}
}

func TestBorrow(t *testing.T) {
	tests := []struct {
		name            string
		available       bool
		wantBorrowed    bool
		wantAvailability bool
	}{
		{
			name:             "available book can be borrowed",
			available:        true,
			wantBorrowed:     true,
			wantAvailability: false,
		},
		{
			name:             "borrowed book cannot be borrowed again",
			available:        false,
			wantBorrowed:     false,
			wantAvailability: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			book := NewBook(1, "Book", "Author", 2024, test.available)
			borrowed := book.Borrow()

			if borrowed != test.wantBorrowed {
				t.Fatalf("Borrow() = %t, want %t", borrowed, test.wantBorrowed)
			}
			if book.Available != test.wantAvailability {
				t.Fatalf("book.Available = %t, want %t", book.Available, test.wantAvailability)
			}
		})
	}
}
