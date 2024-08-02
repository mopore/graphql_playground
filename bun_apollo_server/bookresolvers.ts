import type { Book, BookInput } from "./types";

const books: Book[] = [
    {
        id: "1",
        title: "The Awakening",
        author: "Kate Chopin",
    },
    {
        id: "2",
        title: "City of Glass",
        author: "Paul Auster",
    },
];

export const queryBooks = (): Book[] => books;

export const addBook = (_: any, args: any): Book[] => {

    const bookInput = args.bookInput as BookInput;

    const newBook = {
        id: String(books.length + 1),
        title: bookInput.title,
        author: bookInput.author,
    }
    books.push(newBook);
    console.log(`Added book w/ bookInput: ${newBook.title} by ${newBook.author}`);
    console.log(`New length: ${books.length}`);
    return books;
}


