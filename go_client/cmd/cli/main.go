package main

import (
    "context"
    "fmt"
    "log"

    "github.com/machinebox/graphql"
)

type Book struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}


func main() {
    client := graphql.NewClient("http://localhost:4000/graphql")

    req := graphql.NewRequest(`
        query GetBooks {
          books {
            title
            author
          }
        }
    `)

    var respDate struct {
        Books []Book `json:"books"`
    }

    if err := client.Run(context.Background(), req, &respDate); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Books:")
    for _, book := range respDate.Books {
        fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
    }
}
