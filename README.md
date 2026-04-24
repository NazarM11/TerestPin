# TerestPin Backend 📍

TerestPin is a high-performance, modular RESTful API built with **Go**, designed for an image-sharing platform. The project demonstrates professional backend engineering practices, including **Clean Architecture**, **Dependency Injection**, and type-safe database interactions.

---

## 🛠 Tech Stack

* **Language:** Go (1.23+)
* **Routing:** [Chi Router](https://github.com/go-chi/chi) (Lightweight & idiomatic)
* **Database:** PostgreSQL with [sqlc](https://sqlc.dev/) (Type-safe SQL)
* **Authentication:** JWT (JSON Web Tokens) with Argon2id password hashing
* **Containerization:** Docker & Docker Compose
* **Validation:** Custom logic for file types and metadata

---

## 🏗 Architectural Overview

The project is structured to ensure high maintainability and separation of concerns by following **Dependency Injection** principles.

* **`internal/api`**: Core configurations and shared types (e.g., `ApiConfig`).
* **`internal/api/handlers`**: Pure HTTP logic. Handlers are implemented as closure-based constructors to receive dependencies without global state.
* **`internal/api/middleware`**: Security and logging layers (JWT validation, request logging).
* **`internal/database`**: Auto-generated type-safe DAO layer using `sqlc`.
* **`internal/auth`**: Low-level security primitives for token generation and validation.



---

## 🚀 Key Features

### 🔐 Secure Authentication
* State-of-the-art **Argon2id** hashing for user passwords.
* Secure JWT issuance and validation.
* Custom **Auth Middleware** that injects user identity into the request context.

### 🖼 Media Management
* **Smart Uploads:** Validates file signatures (magic bytes) to ensure uploaded files are actual images, not just renamed binaries.
* **Automated Storage:** Structured file naming using UUIDs to prevent filename collisions.
* **Size Constraints:** Multi-part form limits to prevent DOS attacks via large file uploads.

### 📊 Database Efficiency
* Type-safe queries via `sqlc`, eliminating "magic strings" in SQL.
* Structured migrations for consistent schema management.

---

## 🛠 Getting Started

### Prerequisites
* Go 1.23 or higher
* Docker and Docker Compose
* A `.env` file (see configuration)

### Installation & Run

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/NazarM11/TerestPin.git](https://github.com/NazarM11/TerestPin.git)
    cd TerestPin
    ```

2.  **Configure Environment:**
    Create a `.env` file in the root directory:
    ```env
    # --- Database Configuration ---
    # Used by Docker to initialize the DB and by Go to connect
    DB_USER=postgres
    DB_PASSWORD=your_secure_password_here
    DB_NAME=terestpin_db
    DB_HOST=db
    DB_PORT=5432

    # Connection string for the Go application
    # Format: postgres://<user>:<password>@<host>:<port>/<dbname>?sslmode=disable
    DB_URL=postgres://postgres:your_secure_password_here@db:5432/terestpin_db?sslmode=disable

    # --- Server Configuration ---
    PORT=8080
    UPLOAD_PATH="./uploads"

    # --- Security ---
    # Change this to a long, random string in production
    JWT_SECRET="your_secret_key_here"
    ```

3.  **Spin up the stack:**
    ```bash
    docker-compose up --build
    ```

The server will be live at `http://localhost:8080`.

---

## 🗺 API Endpoints

### Public
* `POST /api/users` - Register a new user
* `POST /api/login` - Authenticate and receive a token
* `GET /api/pins` - List all public pins

### Protected (Requires Bearer Token)
* `POST /api/pins` - Upload a new pin (Multipart Form)
* `DELETE /api/pins/{id}` - Remove a pin
* `DELETE /api/users` - Delete user account and associated data

---

## 🧠 What I Learned
* **Refactoring from Monolith to Modules:** Transitioned the codebase from a single `main.go` to a professional package-based structure.
* **Context Management:** Effectively using `context.Context` to pass user metadata through middleware to handlers.
* **Dependency Injection:** Avoiding global variables by passing configuration pointers to constructors.