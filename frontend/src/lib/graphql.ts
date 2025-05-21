import { GraphQLClient } from 'graphql-request';

export const graphqlClient = new GraphQLClient('http://localhost:8080/query', {
  // Add any headers if needed
  headers: {
    'Content-Type': 'application/json',
  },
});