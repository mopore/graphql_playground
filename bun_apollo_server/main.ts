import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";

const schemaDef = `#graphql
    # Just a comment
    
    type Book {
        id: ID!
        title: String!
        author: String!
    }

    type Query {
        books: [Book]
    }

    type Mutation {
        addBook(title: String!, author: String!): [Book]
    }
`;

interface Book {
    id: string;
    title: string;
    author: string;
}


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

const resolvers = {
    Query: {
        books: (): Book[] => books,
    },
    Mutation: {
        addBook: (_: any, args: any): Book[] => {
            const newBook = {
                id: String(books.length + 1),
                title: args.title,
                author: args.author,
                }
            books.push(newBook);
            console.log(`Added book: ${newBook.title} by ${newBook.author}`);
            console.log(`New length: ${books.length}`);
            return books;
        }
    }
};

const server = new ApolloServer({ typeDefs: schemaDef, resolvers });

const { url } = await startStandaloneServer(
    server, 
    {
        listen: { 
            port: 4000
        } 
    }
);

console.log(`🚀 Server ready at ${url}`);
