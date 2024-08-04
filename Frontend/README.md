# Frontend for Blooogy

The frontend for the Blooogy blog is built with Next.js 14, using NextAuth v5 for authentication and `jwt-decode` for handling JWT tokens. This application utilizes Yarn as the package manager.

## Packages Used

- `next` – React framework for building server-side rendered applications.
- `next-auth` – Authentication library for Next.js.
- `jwt-decode` – Library for decoding JWT tokens.
- `yarn` – Package manager for managing dependencies.

## Setup and Running Locally

To get started, follow these steps:

1. **Install Dependencies**

   Run the following command to install all dependencies:

   ```bash
   yarn install
   ```

2. **Run the Development Server**

   Start the development server with:

   ```bash
   yarn dev
   ```

## Environment Variables

Ensure the following environment variables are set for local development:

### AUTH

- `AUTH_SECRET`: Secret key for authentication. Example: `BX7Xl+EOLvqH2/GajGtUYrdF9bJE7TlqG+2SVmimAUM=`
- `AUTH_URL`: URL for authentication. Example: `http://localhost:3000`

### API

#### UTILS

- `NEXT_PUBLIC_BASE_API`: Base API URL for the backend. Example: `http://backend:4000/api`

#### Auth

- `NEXT_PUBLIC_LOGIN`: Endpoint for login. Example: `/auth/login`
- `NEXT_PUBLIC_REGISTER`: Endpoint for registration. Example: `/auth/register`

#### GETTERS

- **Posts**

  - `NEXT_PUBLIC_GET_ALL_POSTS`: Endpoint to get all posts. Example: `/posts`
  - `NEXT_PUBLIC_GET_SINGLE_POST`: Endpoint to get a single post by ID. Example: `/posts/`
  - `NEXT_PUBLIC_GET_USER_ALL_POSTS`: Endpoint to get all posts by a user. Example: `/posts/user`

- **Users**
  - `NEXT_PUBLIC_GET_ALL_USERS`: Endpoint to get all users. Example: `/users`
  - `NEXT_PUBLIC_GET_SINGLE_USER`: Endpoint to get a single user by ID. Example: `/users/`

#### SETTERS

- **Posts**

  - `NEXT_PUBLIC_EDIT_POST`: Endpoint to edit a post by ID. Example: `/posts/`
  - `NEXT_PUBLIC_DELETE_POST`: Endpoint to delete a post by ID. Example: `/posts/`

- **Users**
  - `NEXT_PUBLIC_DELETE_USER`: Endpoint to delete a user by ID. Example: `/users/`

## Troubleshooting

**Frontend Issues**: Ensure that the backend service is running and accessible at the `NEXT_PUBLIC_BASE_API` URL. Check the console for any errors related to API calls or authentication.

## Docker Deployment

For Docker deployment, you can use the Docker Compose files for building and running the frontend service.

**Development**:

```bash
docker-compose -f docker-compose.dev.yaml up --build
```

**Production**:

```bash
docker-compose up --build
```

**Note:** Ensure that all environment variables are correctly set in your Docker configuration files.

## API Integration

After setting up the frontend, ensure that API endpoints match those defined in the backend. Use tools like Postman or curl for testing endpoints if needed.
