package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func queryBooks(client *graphql.Client) []Book {
	req := graphql.NewRequest(`
        query GetBooks {
            books {
                id
                title
                author
            }
        }
    `)

	var respData struct {
		Books []Book `json:"books"`
	}

	if err := client.Run(context.Background(), req, &respData); err != nil {
		log.Fatal(err)
	}

    return respData.Books
}

func addBook(client *graphql.Client) []Book {
	req := graphql.NewRequest(`
        mutation AddBook($title: String!, $author: String!) {
            addBook(title: $title, author: $author) {
                id
                title
                author
            }
        }
    `)
	req.Var("title", "The Art of Computer Programming")
	req.Var("author", "Donald E. Knuth")

	var respData struct {
		Books []Book `json:"addBook"`
	}

	if err := client.Run(context.Background(), req, &respData); err != nil {
		log.Fatal(err)
	}

    return respData.Books
}

func printBooks(books []Book) {
    fmt.Println("Books:")
    fmt.Printf(("Length of books: %d\n"), len(books))
    for _, book := range books {
        fmt.Printf("%s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}


func main() {
	client := graphql.NewClient("http://localhost:4000/graphql")

    fmt.Println("Querying books...")
    books := queryBooks(client)
    printBooks(books)

    fmt.Println("Adding a new book...")
    books = addBook(client)
    printBooks(books)
}
