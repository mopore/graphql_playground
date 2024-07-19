import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";

const schemaDef = `#graphql
    # Just a comment
    
    type Book {
        title: String
        author: String
    }

    type Query {
        books: [Book]
    }

`;

const books = [
    {
        title: "The Awakening",
        author: "Kate Chopin",
    },
    {
        title: "City of Glass",
        author: "Paul Auster",
    },
];

const resolvers = {
    Query: {
        books: () => books,
    },
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
