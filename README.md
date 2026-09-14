# Snip-It

A lightweight code/text snippet sharing web application built with Go and PostgreSQL.

![Snip-It Screenshot](ui/static/img/Screenshot.png)

## Overview

Snip-It lets users create, store, and view text snippets through a simple web interface. It's built with the Go standard library's `net/http`, a PostgreSQL backend via `pgx`, and server-rendered HTML templates — no heavy frameworks involved.

## Features

- Create and view text snippets
- Server-side HTML templating with a template cache for performance
- PostgreSQL persistence via `pgx/v5` connection pooling
- Environment-based configuration with `.env` support
- Static asset serving (CSS/JS) via a dedicated file server

## Tech Stack

- **Language:** Go 1.26+
- **Database:** PostgreSQL (via [pgx/v5](https://github.com/jackc/pgx))
- **Config:** [godotenv](https://github.com/joho/godotenv)
- **Static files:** [fileonlyserver](https://github.com/seanzhengw/fileonlyserver)

## Project Structure

```
snip-it/
├── handlers/     # HTTP handlers and template cache
├── internal/     # Application dependencies, config loading, models
├── routes/       # Route registration
├── ui/           # HTML templates and static assets
├── main.go       # Application entry point
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Go 1.26 or later
- A running PostgreSQL instance

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mohnasr137/snip-it.git
   cd snip-it
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the project root:
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/snipit?sslmode=disable
   PORT=:4000
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

5. Open your browser at [http://localhost:4000](http://localhost:4000)

## Configuration

| Variable       | Description                          | Example                                              |
|----------------|---------------------------------------|-------------------------------------------------------|
| `DATABASE_URL` | PostgreSQL connection string          | `postgres://user:pass@localhost:5432/snipit`         |
| `PORT`         | Port the server listens on            | `:4000`                                              |

## Contributing

Contributions, issues, and feature requests are welcome. Feel free to open a pull request or issue.

## License

This project currently has no license specified. Add one (e.g. MIT) if you plan to open it up for reuse.
