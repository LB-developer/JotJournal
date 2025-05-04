# JotJournal

JotJournal’s core goal is simplicity: to provide a lightweight, no-nonsense way to annotate and track habits ("jots") without unnecessary complexity or distractions. This focus on minimalism addresses a common pain point in habit tracking - many apps overcomplicate what should be effortless.

---

## Live API & Documentation

The API is currently deployed and accessible online. Interactive Swagger documentation is available at:  
https://jotjournal.cloud/swagger/index.html#

This Swagger UI provides a user-friendly interface to explore and test all available endpoints.

---

## Tech Stack & Architecture

- **Backend:** Go 1.23.2, RESTful API, JWT authentication
- **DevOps:** GitHub Actions, AWS Codepipeline  
- **Testing:** Dockertest, Table-driven testing

---

## Highlights

- Clean, modular Go API with focus on maintainability  
- CI/CD pipeline:  
  - Automated build & testing via GitHub Actions
  - Production deployment is triggered automatically via AWS CodePipeline when changes are merged to `main`
- Secure JWT-based authentication and password handling  
- Dockerized for consistent local and cloud deployment  
- Swagger/OpenAPI documentation for easy API exploration and integration  

---

## Quick Start

### Local Mode (Go & PostgreSQL)
Use this if you have Go 1.23+ installed and a local Postgres instance.

1. Clone the repository:  
       `git clone https://github.com/LB-developer/JotJournal.git && cd JotJournal`

2. Configure environment variables:  
       `cp server/.env.example server/.env`

3. Apply database migrations:  
       `make -C server migrate-up`

4. Start the API server:  
       `make -C server run`

The server will listen on port 8080 of your localhost and connect to your local database.

---

### Container Mode (Docker & Docker Compose)
Use this if you prefer a containerized setup without installing Go or Postgres.

1. Clone the repository:  
       `git clone https://github.com/LB-developer/JotJournal.git && cd JotJournal`

2. Configure environment variables:  
       `cp server/.env.example server/.env`

3. Run everything with one command:  
       `docker compose --profile dev up --build`

This will:
- Spin up a Postgres container  
- Run migrations automatically  
- Launch the API on port 8080  

---

### View the API  
Open in your browser:  
       `http://localhost:8080/swagger/index.html#/` 
to explore the interactive Swagger UI and test endpoints.

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
