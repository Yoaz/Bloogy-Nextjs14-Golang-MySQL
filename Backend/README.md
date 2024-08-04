# Backend for Blooogy

The backend for the Blooogy blog is built with Golang, using GORM for ORM, Chi for routing, and MySQL as the database. The application is designed to run inside a Docker container for both development and production environments.

## Packages Used

- `godotenv` – For loading environment variables from a `.env` file.
- `chi` – A lightweight, idiomatic router for Go.
- `bcrypt` – For hashing passwords.
- `mysql` – MySQL driver for Go.
- `gorm` – ORM library for Go.
- `cors` – For handling Cross-Origin Resource Sharing.

## Setup and Running Locally

To get started, follow these steps:

1. **Run Locally**

   Use the provided `Makefile` to start the application:

   ```bash
   make run
   ```

## Environment Variables

Ensure the following environment variables are set for local development:

### MySQL

- `MYSQL_ROOT_PASSWORD=<your_mysql_root_password>`
- `DB_USER=<your_mysql_user>`
- `DB_PASSWORD=<your_mysql_password>`
- `DB_HOST=127.0.0.1`
- `DB_PORT=<your_db_port>`
- `DB_NAME=<your_db_name>`

### Config

- `PORT=<your_port>`
- `HOST=<your_host_address>`

### Auth

- `API_SECRET=<your_auth_api_secret>`

## Database Initialization

The application will automatically create the MySQL database if it does not already exist. Ensure that MySQL is installed and running on your machine.

## Troubleshooting

**Backend Service Issues**: Ensure that the database service is fully established before starting the backend service.

## Docker Deployment

### Docker Compose for Development

For development with auto-reload and other conveniences, use the development Docker Compose file:

```bash
docker-compose -f docker-compose.dev.yaml up --build
```

### Docker Compose for Production

For production deployment, use the main Docker Compose file:

```bash
docker-compose up --build
```

**Note:** Ensure that the database service has fully started before attempting to start the backend service.

## API Integration

After deploying, use tools like Postman or curl to create an Admin user. Use the payload with **"role": "admin"** (default role is **"user"**).
