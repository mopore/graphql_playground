package main

import (
    "context"
    "fmt"
    "log"

    "github.com/machinebox/graphql"
)

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
        Books []struct {
            Title  string `json:"title"`
            Author string `json:"author"`
        } `json:"books"`
    }

    if err := client.Run(context.Background(), req, &respDate); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Books:")
    for _, book := range respDate.Books {
        fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
    }
}
