# E-STORE

A full-stack e-commerce web application built to learn backend development, REST APIs, and real-world Git workflows.

## Tech Stack

| Layer | Technology |
|---|---|
| Frontend | React |
| Backend | Go (Golang) |
| Database | PostgreSQL |
| Version Control | Git + GitHub |

## Branch Strategy

| Branch | Purpose |
|---|---|
| `main` | Production — stable, never commit directly here |
| `dev` | Integration — all features merge here first |
| `feature/xxx` | Where all new code is written day-to-day |

## Project Structure

```
E-STORE/
├── backend/
│   ├── main.go          # Entry point — starts the server
│   ├── handlers/        # API route handlers (what happens when an endpoint is hit)
│   ├── models/          # Database models (Product, User, Order etc.)
│   └── db/              # Database connection setup
├── frontend/            # React app
└── README.md
```

## Getting Started

### Backend
```bash
cd backend
go run main.go
```

### Frontend
```bash
cd frontend
npm install
npm start
```

## Learning Goals
- How REST APIs work
- How to test API endpoints
- How to raise and review Pull Requests
- How backend and frontend connect
