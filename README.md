🎯 Motivation

TerestPin was built to deepen my understanding of production-grade backend systems in Go. The goal was to move beyond small exercises and design a real-world REST API with clear boundaries between components, strong typing at the database layer, and secure authentication flows.

This project specifically focuses on:

Building a scalable and maintainable backend architecture
Practicing clean separation of concerns using Clean Architecture principles
Implementing secure authentication with JWT and modern password hashing (Argon2id)
Working with type-safe SQL generation using sqlc
Simulating real-world concerns like file validation, upload limits, and structured storage
⚡ Quick Start

If you want to run the project locally in under a minute:

git clone https://github.com/NazarM11/TerestPin.git
cd TerestPin
cp .env.example .env   # or manually create .env as shown above
docker-compose up --build

Then visit:

http://localhost:8080
📖 Usage
Register a user
curl -X POST http://localhost:8080/api/users \
-H "Content-Type: application/json" \
-d '{"email":"test@example.com","password":"password123"}'
Login
curl -X POST http://localhost:8080/api/login \
-H "Content-Type: application/json" \
-d '{"email":"test@example.com","password":"password123"}'

Response will include a JWT token:

{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
Create a pin (protected)
curl -X POST http://localhost:8080/api/pins \
-H "Authorization: Bearer YOUR_TOKEN" \
-F "image=@example.jpg"
Get pins
curl http://localhost:8080/api/pins
🤝 Contributing

Contributions are welcome, but this is primarily a learning-focused project.

If you want to contribute:

Fork the repository

Create a feature branch:

git checkout -b feature-name
Make your changes with clear, atomic commits
Ensure the project builds and runs via Docker
Open a pull request describing what you changed and why
Guidelines
Keep code consistent with Clean Architecture principles used in the project
Avoid introducing global state
Prefer explicit dependencies over hidden coupling
Keep handlers thin and business logic separated
