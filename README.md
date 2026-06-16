TerestPin Backend 📍

TerestPin is a high-performance, modular RESTful API built with Go, designed for an image-sharing platform. The project demonstrates professional backend engineering practices, including Clean Architecture, Dependency Injection, and type-safe database interactions.

💡 Motivation

Most modern image-sharing applications face complex challenges around asset security, rapid database lookups, and scaling heavy file uploads. TerestPin was built to explore how to solve these problems using Go's standard library philosophy: low overhead, absolute clarity, and no magical abstraction layers.

The goal of this project was to move away from massive, unmaintainable monolithic files and build an enterprise-ready API structure from scratch. It serves as a blueprint for implementing strict compile-time safety (via sqlc), bulletproof security parameters (via Argon2id), and highly modular dependency injection without relying on bloated third-party frameworks.

🛠 Tech Stack

    Language: Go (1.23+)

    Routing: Chi Router (Lightweight & idiomatic)

    Database: PostgreSQL with sqlc (Type-safe SQL)

    Authentication: JWT (JSON Web Tokens) with Argon2id password hashing

    Containerization: Docker & Docker Compose

    Validation: Custom logic for file types and metadata

🏗 Architectural Overview

The project is structured to ensure high maintainability and separation of concerns by following Dependency Injection principles.

    internal/api: Core configurations and shared types (e.g., ApiConfig).

    internal/api/handlers: Pure HTTP logic. Handlers are implemented as closure-based constructors to receive dependencies without global state.

    internal/api/middleware: Security and logging layers (JWT validation, request logging).

    internal/database: Auto-generated type-safe DAO layer using sqlc.

    internal/auth: Low-level security primitives for token generation and validation.

🚀 Key Features

🔐 Secure Authentication

    State-of-the-art Argon2id hashing for user passwords.

    Secure JWT issuance and validation.

    Custom Auth Middleware that injects user identity into the request context.

🖼 Media Management

    Smart Uploads: Validates file signatures (magic bytes) to ensure uploaded files are actual images, not just renamed binaries.

    Automated Storage: Structured file naming using UUIDs to prevent filename collisions.

    Size Constraints: Multi-part form limits to prevent DOS attacks via large file uploads.

📊 Database Efficiency

    Type-safe queries via sqlc, eliminating "magic strings" in SQL.

    Structured migrations for consistent schema management.

⚡ Quick Start

For those who want to see the application running instantly with default configurations:
Bash

# 1. Clone and enter the project
git clone https://github.com/NazarM11/TerestPin.git && cd TerestPin

# 2. Copy the default environment template
cp .env.example .env

# 3. Launch the environment
docker-compose up --build

The server will boot up and begin listening for traffic at http://localhost:8080.
🛠 Getting Started
Prerequisites

    Go 1.23 or higher

    Docker and Docker Compose

    A .env file (see configuration)

Installation & Run

    Clone the repository:
    Bash

    git clone https://github.com/NazarM11/TerestPin.git
    cd TerestPin

    Configure Environment:
    Create a .env file in the root directory:
    Code snippet

    # --- Database Configuration ---
    DB_USER=postgres
    DB_PASSWORD=your_secure_password_here
    DB_NAME=terestpin_db
    DB_HOST=db
    DB_PORT=5432

    # Connection string for the Go application
    DB_URL=postgres://postgres:your_secure_password_here@db:5432/terestpin_db?sslmode=disable

    # --- Server Configuration ---
    PORT=8080
    UPLOAD_PATH="./uploads"

    # --- Security ---
    JWT_SECRET="your_secret_key_here"

    Spin up the stack:
    Bash

    docker-compose up --build

🕹 Usage

Once the server is up and running, you can interact with the API endpoints using tools like curl, Postman, or Bruno.
1. Register a New User
Bash

curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username": "gopher123", "email": "gopher@example.com", "password": "supersecurepassword"}'

2. Log In to Receive a JWT Bearer Token
Bash

curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email": "gopher@example.com", "password": "supersecurepassword"}'

    Note: Save the "token" value returned in the JSON response. You will need to pass it in the authorization header for protected actions.

3. Upload a Image Pin (Protected)
Bash

curl -X POST http://localhost:8080/api/pins \
  -H "Authorization: Bearer <YOUR_JWT_TOKEN>" \
  -F "title=My First Pin" \
  -F "image=@/path/to/your/image.jpg"

🗺 API Endpoints
Public

    POST /api/users - Register a new user

    POST /api/login - Authenticate and receive a token

    GET /api/pins - List all public pins

Protected (Requires Bearer Token)

    POST /api/pins - Upload a new pin (Multipart Form)

    DELETE /api/pins/{id} - Remove a pin

    DELETE /api/users - Delete user account and associated data

🧠 What I Learned

    Refactoring from Monolith to Modules: Transitioned the codebase from a single main.go to a professional package-based structure.

    Context Management: Effectively using context.Context to pass user metadata through middleware to handlers.

    Dependency Injection: Avoiding global variables by passing configuration pointers to constructors.

🤝 Contributing

Contributions are welcome! Please follow these simple steps to contribute to the project:

    Fork the repository.

    Create a new feature branch (git checkout -b feature/amazing-feature).

    Commit your changes with clear, descriptive messages (git commit -m 'Add amazing new feature').

    Push your branch to GitHub (git push origin feature/amazing-feature).

    Open a Pull Request against the main branch.
