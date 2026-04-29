### Description ###
This project is a simple Book Management REST API built in Go, demonstrating basic CRUD operations with a PostgreSQL database using the GORM ORM and Gorilla Mux router. It allows users to create, read, update, and delete book records through HTTP endpoints, serving JSON responses. The project is structured with controllers, models, routes, and utility packages, and is currently being refactored toward a cleaner architecture by introducing a service layer to separate business logic from HTTP handling. It serves as a learning project for understanding backend development in Go, API design, and database integration.

## How to run ##
1. go mod tidy
2. go run cmd/main/main.go
3. http://localhost:9010
