import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";
import { addBook, queryBooks } from "./bookresolvers";


const SCHEMA_DEF_PATH = "./schema.graphql" as const;


const mainAsync = async () => {
    const schemaDef = await Bun.file(SCHEMA_DEF_PATH).text();

    const resolvers = {
        Query: {
            books: queryBooks,
        },
        Mutation: {
            addBook: addBook,        
        }
    };

    const server = new ApolloServer({ 
        typeDefs: schemaDef, 
        resolvers 
    });

    const { url } = await startStandaloneServer(
        server, 
        {
            listen: { 
                port: 4000
            } 
        }
    );

    console.log(`🚀 Server ready at ${url}`);
}

await mainAsync();

