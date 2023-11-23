## Structure
```
.
├── README.md
├── go.mod
├── go.sum
├── internal
│   ├── core
│   │   ├── domain
│   │   │   └── message.go
│   │   ├── ports
│   │   │   ├── repositories
│   │   │   │   └── message.go
│   │   │   └── services
│   │   │       └── message.go
│   │   └── services
│   │       └── message.go
│   ├── handlers
│   │   ├── message.go
│   │   └── model
│   │       └── message.go
│   ├── infrastructure
│   │   └── datastore
│   │       └── oracle-datastore.go
│   └── repositories
│       ├── message.go
│       └── model
│           └── message.go
├── main.go
└── pkg
    └── config
        └── config.go
```

- pkg: utils, shared function etc.
- core: keep everything related to business logic.
- infrastructure: database initiation, routing etc.
- handlers: handling http request.
- repositories: involved in querying data and making API calls to external services.

## Concept
Ports and Adapters

Anything within the core domain containing our business logic should be unaware of external details. It shouldn't be aware of the database or framework used, focusing solely on business logic. External components communicate with it through the provided core interfaces.

## Flow
After user makes an HTTP request:
1. The request is routed to the handler, which call services inside core via the interface
2. Services handle business logic and call repository via the interface for data from database or external api services.

## Pros
- Easy to change: For example: changing the database won't affect the business logic.
- Seperation of concerns.

## Cons
- Overkill for small projects.
- Takes more time to start the project.

## Remark
- This structure is too complicate for small and simple project.
- For a simple project we can have only handler and services
    - handler: handles http requests and logic.
    - services: performs queries.
