# Login-Go API

A RESTful API built with **Go** and **JWT** for authentication and authorization. This project provides a simple login system using JWT tokens to protect routes and access user resources.

## Description

This project demonstrates how to create an API in **Go** that uses JSON Web Tokens (JWT) to protect routes and allow access only to authenticated users.

### Features

- **JWT Authentication**: Users can authenticate and receive a JWT token.
- **Protected Routes**: Only users with a valid JWT token can access protected routes.
- **Modular Structure**: The project is structured into separate packages to keep the code clean and organized.

## Requirements

- **Go 1.18+**
- **Postman** or any API client for testing the routes.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/JunniorRavelo/login-go.git
   cd login-go
   ```

2. Install the dependencies:

   ```bash
   go mod tidy
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

   The server will be running at `http://localhost:8080`.

## Routes

### 1. **Login (POST /login)**

Generates a JWT token for the user.

**Request:**

- URL: `/login`
- Method: `POST`
- Body (JSON):
  
  ```json
  {
    "username": "admin",
    "password": "password"
  }
  ```

**Response:**

If the credentials are correct, you will receive the JWT token in the response:

```json
{
  "token": "your_jwt_token_here"
}
```

### 2. **Get User (GET /user)**

Fetches the user details (requires a valid JWT token).

**Request:**

- URL: `/user`
- Method: `GET`
- Headers:

  | Key           | Value                                  |
  |---------------|----------------------------------------|
  | Authorization | Bearer <your_jwt_token_here>           |

**Response:**

```json
{
  "username": "admin",
  "email": "admin@example.com"
}
```

## Testing

You can test the API using **Postman** or any HTTP client with the following steps:

1. **Login**:
   - Send a `POST` request to `/login` with the user credentials.
   - You will receive a JWT token in the response.

2. **Access Protected Routes**:
   - Use the token received in the previous step to send requests to protected routes.
   - Add the token to the `Authorization: Bearer <your_jwt_token_here>` header.

## Project Structure

```
login-go/
│
├── main.go          # Main entry point
├── handlers/        # Handlers to manage routes
│   ├── auth.go      # Handles login and JWT generation
│   ├── user.go      # Example protected route
│
├── middleware/      # Middleware to verify JWT
│   └── jwt.go       # JWT verification and validation
│
├── models/          # Data models
│   └── user.go      # User structure
│
├── utils/           # Utilities, like JWT generation
│   └── jwt.go       # Functions to generate and validate JWT
│
├── go.mod           # Go module
└── go.sum           # Dependencies
```

## Contributing

If you'd like to contribute to this project, follow these steps:

1. Fork this repository.
2. Create a new branch (`git checkout -b feature/new-feature`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push the branch (`git push origin feature/new-feature`).
5. Create a pull request.

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more details.