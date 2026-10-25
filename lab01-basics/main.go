package main

import "fmt"

const targetAuthor = "George Orwell"

func main() {
	books := []Book{
		NewBook(1, "1984", "George Orwell", 1949, true),
		NewBook(2, "Animal Farm", "George Orwell", 1945, false),
		NewBook(3, "The Go Programming Language", "Alan Donovan", 2015, true),
		NewBook(4, "Effective Java", "Joshua Bloch", 2018, true),
	}

	fmt.Println("All books:")
	for _, book := range books {
		fmt.Println(book.Description())
	}

	fmt.Printf("\nBooks by %s:\n", targetAuthor)
	authorBooks := FindBooksByAuthor(books, targetAuthor)
	switch len(authorBooks) {
	case 0:
		fmt.Println("No books found")
	case 1:
		fmt.Println("One book found")
	default:
		fmt.Printf("%d books found\n", len(authorBooks))
	}

	for _, book := range authorBooks {
		fmt.Println(book.Description())
	}

	fmt.Printf("\nAvailable books: %d\n", CountAvailableBooks(books))

	newestBook, found := FindNewestBook(books)
	if found {
		fmt.Println("Newest book:", newestBook.Description())
	}

	fmt.Println("\nBorrowing the first book:")
	if books[0].Borrow() {
		fmt.Println("The book was borrowed successfully")
	} else {
		fmt.Println("The book is already borrowed")
	}
	fmt.Println(books[0].Description())

	fmt.Println("\nBooks grouped by author:")
	groups := GroupBooksByAuthor(books)
	for author, authorBooks := range groups {
		fmt.Printf("%s: %d\n", author, len(authorBooks))
	}
}
