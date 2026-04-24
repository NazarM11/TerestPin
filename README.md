# Terestpin Backend

A robust, production-ready backend service inspired by Pinterest, built with **Go (Golang)**, **PostgreSQL**, and **Docker**. This project demonstrates clean architecture, JWT authentication, and automated environment management.

## 🚀 Features

- **User Management**: Secure registration and authentication.
- **JWT Authentication**: Protected routes using JSON Web Tokens.
- **RESTful API**: Clean and predictable API structure.
- **PostgreSQL Integration**: Persistent storage for users and pins.
- **Automated Cleanup**: Logic for synchronizing database records with local file storage.
- **Dockerized Infrastructure**: Fully containerized setup with health checks for reliable startup.

## 🛠 Tech Stack

- **Language**: Go 1.23
- **Database**: PostgreSQL 15
- **Router**: Chi 
- **Containerization**: Docker & Docker Compose
- **Environment**: Dotenv for secure secret management

## 📦 Project Structure

```text
├── backend/
│   ├── main.go          # Entry point
│   ├── internal/        # Logic & handlers
│   ├── uploads/         # Local image storage (ignored by git)
│   └── Dockerfile       # Multi-stage build for the backend
├── docker-compose.yml   # Infrastructure orchestration
├── .env.example         # Template for environment variables
└── .gitignore           # Safety first!
