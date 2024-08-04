# Blooogy

## Overview

Blooogy is a full-stack blog application built with **Next.js 14**, **Golang**, **MySQL**, and **Docker**. This guide provides instructions for setting up the project in both development and production environments.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
  - [Production Setup](#production-setup)
  - [Development Setup](#development-setup)
- [Creating an Admin User](#creating-an-admin-user)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## Features

- **Backend**: Golang-based API
- **Frontend**: Next.js 14 with server-side rendering
- **Database**: MySQL
- **Deployment**: Docker for containerization

## Prerequisites

- Docker
- Docker Compose

## Setup

### Production Setup

To run the application in production mode:

1. Build and start the services with Docker Compose:

   ```bash
   docker-compose up --build
   ```

2. This command will build the Docker images and start the containers for the backend, frontend, and MySQL services.

### Development Setup

For development purposes, you can use a separate Docker Compose configuration to support features like auto-reloading:

1. Ensure you have a `docker-compose.dev.yaml` file in the root directory.

2. Start the development services:

   ```bash
   docker-compose -f docker-compose.dev.yaml up --build
   ```

3. This setup may include live-reload features and other development conveniences.

**Note:** If the backend service fails to start, wait for the database service to be fully established and then restart the backend service.

## Creating an Admin User

To create an Admin user, use Postman or `curl` to send a request to the `/auth/register` endpoint with the following payload:

```json
{
  "email": "adminUser",
  "password": "adminPassword",
  "role": "admin" // Default role is "user"
}
```

TODO:

- Fix rendering of dashboard in frontend once adding user/editing post
- Fix uncaught error for forms: Add New User, Login, Register
- Fix register success not routing to homepage
- Hash password in fronend before sending backend
- Fix on mobile view - scrolling right extra space
- Finish contact form backend send email
- Add edit user option in frontend
- More styling for frontend
