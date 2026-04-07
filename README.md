# Taskly Backend

Taskly Backend is a robust RESTful API built with **Go** to power a task management application. It leverages the **Gin** web framework for high-performance routing and **GORM** for seamless PostgreSQL database interactions and migrations.

## 🚀 Features

  - **RESTful API**: Clean and scalable endpoints for task management.
  - **Automated Migrations**: Uses GORM's `AutoMigrate` to keep your database schema in sync with your Go models.
  - **High Performance**: Built with the Gin framework for minimal overhead and fast request processing.
  - **Database Support**: Native integration with PostgreSQL.

## 🛠️ Tech Stack

  - **Language**: [Go (Golang)](https://golang.org/)
  - **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
  - **ORM**: [GORM](https://gorm.io/)
  - **Database**: [PostgreSQL](https://www.postgresql.org/)

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

  - Go (1.20 or later)
  - PostgreSQL (running locally or via Docker)
  - Git

## ⚙️ Installation & Setup

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/sriram32005/taskly-backend.git
    cd taskly-backend
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Configure the Database:**
    Ensure your PostgreSQL service is running and create a database named `taskly`. Update the DSN (Data Source Name) in `main.go` or your configuration file:

    ```go
    dsn := "host=localhost user=postgres password=yourpassword dbname=taskly port=5432 sslmode=disable"
    ```

4.  **Run the application:**

    ```bash
    go run main.go
    ```

## 🏗️ Project Structure

```text
taskly-backend/
├── main.go          # Application entry point and database initialization
├── models/          # GORM database models
├── routes/          # API route definitions
├── controllers/     # Business logic and request handling
├── vendor/          # Project dependencies (if vendored)
└── go.mod           # Go module file
```
