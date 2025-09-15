# JotJournal

JotJournal’s core goal is to provide a lightweight, no-nonsense way to annotate and track habits ("jots") without unnecessary complexity or distractions. This focus on minimalism addresses a common pain point in habit tracking; many apps overcomplicate what should be effortless.

---

## Live API & Documentation

The API is currently deployed and accessible online. Interactive Swagger documentation is available at:  
https://jotjournal.cloud/swagger/index.html#

This Swagger UI provides a user-friendly interface to explore and test all available endpoints.

---

## Tech Stack & Architecture

- **Backend:** Go 1.23.2, Valkey, RESTful API, JWT authentication,
- **CI/CD:** GitHub Actions, AWS Lightsail
- **Testing:** Dockertest, Table-driven testing, Playwright

---

## Highlights

- Clean, modular Go API with focus on maintainability  
- CI/CD pipeline:  
  - Automated build & testing of code via GitHub Actions
  - Automated build & push of Docker images via GitHub Actions
- Secure JWT-based authentication and password handling  
- Session backed authorization using Valkey
- Dockerized for consistent local and cloud deployment  
- Swagger/OpenAPI documentation for easy API exploration and integration  

---

## Quick Start

### Local Mode (Go, PostgreSQL & Valkey)
Use this if you have Go 1.23+, Node.js 20+, and a local PostgreSQL + Valkey setup and want to run locally.

1. Clone the repository:  
       `git clone https://github.com/LB-developer/JotJournal.git && cd JotJournal`

2. Set up environment variables and local database:  
       `make setup`

3. Apply database migrations:  
       `make migrate`

4. Start backend, frontend, and Valkey locally:  
       `make dev`

This will:
- Initialize a postgres database and valkey instance
- Run migrations automatically
- Launch the API on port 8080  
- Launch the frontend on port 3000

---

### Container Mode (Docker & Docker Compose)
Use this if you prefer a containerized setup without installing Go, Node.js, Valkey or PostgreSQL.

1. Clone the repository:  
       `git clone https://github.com/LB-developer/JotJournal.git && cd JotJournal`

3. Run everything with one command:  
       `make docker-dev`

This will:
- Spin up Postgres, Valkey, Frontend & Backend containers  
- Run migrations automatically  
- Launch the API on port 8080  
- Launch the frontend on port 3000

---

### View the App & API

- Frontend:  
       `http://localhost:3000`

- API (Swagger UI):  
       `http://localhost:8080/swagger/index.html#/`

---

## Roadmap

- Frontend client (NextJS) in development  
    - Monthly habits/task dashboard
    - Analytics dashboard  
    - Calendar integrations  
- Public API documentation enhancements  

---

## Contributing

Fork the repo, create a feature branch, and submit a pull request. Ensure tests pass and code style is followed.

---

JotJournal currently offers a robust, well-documented backend API as a foundation for scalable productivity tools, with plans to create client-side handling soon.
