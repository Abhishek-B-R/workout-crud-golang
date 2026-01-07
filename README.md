# Workout CRUD Application (Go, chi, PostgreSQL)

A simple and robust CRUD (Create, Read, Update, Delete) application for managing workouts, built with Go, the [chi](https://github.com/go-chi/chi) HTTP router, and PostgreSQL as its database.

## Features

- **RESTful API** to create, retrieve, update, and delete workouts
- Built using idiomatic Go and [chi](https://github.com/go-chi/chi) for efficient HTTP routing
- PostgreSQL database integration
- Clean and modular structure for easy extension

## Requirements

- [Go 1.18+](https://golang.org/dl/)
- [PostgreSQL 12+](https://www.postgresql.org/download/)
  - (Docker Compose-based PostgreSQL is available for local development, see below)

## Setup

### 1. Clone the repository and install packages
```bash
git clone https://github.com/Abhishek-B-R/workout-crud-golang.git
go mod download
```

### 2. Run Docker Compose to start PostgreSQL
```bash
docker compose up --build
```

### 3. Start the application
```bash
go run main.go
```