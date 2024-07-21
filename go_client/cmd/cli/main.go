package main

import (
    "context"
    "fmt"
    "log"

    "github.com/machinebox/graphql"
)

type Book struct {
    ID    string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

func queryBooks(client *graphql.Client) {
    req := graphql.NewRequest(`
        query GetBooks {
          books {
            id
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
        fmt.Printf("%s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}


func main() {
    client := graphql.NewClient("http://localhost:4000/graphql")
    queryBooks(client)
}
