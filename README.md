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

## Highlights & Best Practices

- Clean, modular Go API with focus on maintainability  
- CI/CD pipeline:  
  - Automated build & testing via GitHub Actions
  - Production deployment is triggered automatically via AWS CodePipeline when changes are merged to `main`
- Secure JWT-based authentication and password handling  
- Dockerized for consistent local and cloud deployment  
- Swagger/OpenAPI documentation for easy API exploration and integration  

---

## Quick Start

- Clone and run the API locally:
```bash
git clone https://github.com/LB-developer/JotJournal.git
cd JotJournal/server
make build
make test
make run
```

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
