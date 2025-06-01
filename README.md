# go-movies-rest

A simple RESTful API for managing a collection of movies, written in Go (Golang). This project demonstrates basic CRUD (Create, Read, Update, Delete) operations using Go's standard library and is ideal for learning or as a starter template for REST APIs.

## Features

- List all movies
- Get details of a specific movie
- Add a new movie
- Update an existing movie
- Delete a movie

## Technologies Used

- [Go (Golang)](https://golang.org/)
- Standard Go HTTP library (no external frameworks)
- JSON for data serialization

## Getting Started

### Prerequisites

- Go (version 1.18 or newer recommended)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/devhimal/go-movies-rest.git
   cd go-movies-rest
   ```

2. **Build the application:**

   ```bash
   go build -o movies-api
   ```

3. **Run the application:**

   ```bash
   ./movies-api
   ```

   By default, the server will start on port `8000`.

### API Endpoints

| Method | Endpoint       | Description              |
| ------ | -------------- | ------------------------ |
| GET    | `/movies`      | List all movies          |
| GET    | `/movies/{id}` | Get a movie by ID        |
| POST   | `/movies`      | Add a new movie          |
| PUT    | `/movies/{id}` | Update an existing movie |
| DELETE | `/movies/{id}` | Delete a movie by ID     |

#### Example: Add a Movie

```bash
curl -X POST http://localhost:8080/movies \
     -H 'Content-Type: application/json' \
     -d '{"title":"Inception","director":"Christopher Nolan","year":2010}'
```

#### Example: List All Movies

```bash
curl http://localhost:8080/movies
```

## Project Structure

```
├── main.go          # Entry point and HTTP handlers
└── README.md
```

## Customization

You can easily extend the project to use a persistent database (e.g., PostgreSQL, MongoDB) or add authentication.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---

Made with ❤️ by [devhimal](https://github.com/devhimal)
