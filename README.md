# go-chat

## About
- cmd/scd: Entrypoint, the main application or service.

- config/config.go: Contains the configuration setup for the application

- graphql: This directory contains several Go files related to GraphQL
  - docs: Might contain documentation or schemas related to GraphQL.
  - query.go: Defines GraphQL queries.
  - type.go: Defines GraphQL types.
  - schema.go: Defines the GraphQL schema.
  - subscription.go: Handles GraphQL subscriptions, which are used for real-time functionality, a common feature in chat applications.
  - mutation.go: Manages GraphQL mutations for creating, updating, or deleting data.
  - resolver.go: Contains resolver functions that handle the business logic for GraphQL queries and mutations.

- internal: Service code
  - dto: Stands for Data Transfer Objects
  - domain: Typically includes domain models and business logic.

- util/uuid.go: A utility file, possibly for generating or handling UUIDs, which are often used for unique identifiers in applications.
- tool/dynamo: This could be a tool or script related to Amazon DynamoDB, suggesting that DynamoDB might be used as a database for the chat application.
