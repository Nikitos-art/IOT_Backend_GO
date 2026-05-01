# IoT Backend Platform (Go)

This project is a backend system built in Go that simulates a lightweight IoT data platform. It evolved from a simple CRUD-based book management API into a more realistic backend architecture supporting authentication, device management, and data ingestion.

The project was developed using a **Vibe Coding** approach — focusing on rapid iteration, incremental improvements, and system-driven design rather than upfront overengineering.

---

## 🚀 Overview

The system allows users to:

- Register and authenticate using JWT
- Create and manage devices under their account
- Generate API keys for devices
- Allow devices to send data securely via an ingestion API
- Store and retrieve device-related data from a PostgreSQL database

---

## 🧱 Architecture

The project follows a layered architecture:

- **Controllers** – Handle HTTP requests and responses
- **Services** – Contain business logic and database operations
- **Models** – Define database structures (GORM models)
- **Middleware** – Handles authentication and request validation
- **Routes** – Defines API endpoints and routing structure
- **Utils** – Helper functions (JWT, API key generation, etc.)

---

## 🔐 Authentication

### User Authentication
- JWT-based authentication
- Protected routes require `Authorization: Bearer <token>`

### Device Authentication
- Each device is assigned a unique API key
- Devices authenticate via `X-API-Key` header when sending data

---

## 📡 Key Features

- User registration & login
- JWT middleware for protected routes
- Device CRUD operations (owned by users)
- Secure API key generation for devices
- Device data ingestion endpoint
- PostgreSQL integration via GORM
- Clean separation of concerns (service-based architecture)

---

## 🧪 Example Flow

1. User logs in → receives JWT
2. User creates device → receives API key
3. Device sends data:

## How to run ##
1. go mod tidy
2. go run cmd/main/main.go
3. http://localhost:9010

## Git ##
1. Git commands must be run from one level up of the root dir

### PERSONAL NOTES ###
1. Putting validation tags directly on your GORM model is discouraged in large projects due to tight coupling