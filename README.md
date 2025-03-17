# URL Shortener

## Overview
This project is a URL shortening service built using Golang. It allows users to convert long URLs into short, easy-to-share links and retrieve the original URLs from shortened ones.

## Features
- Shorten long URLs
- Retrieve original URLs from short codes
- Rate limiting for API requests
- Authentication for secure access
- Database storage for persistence

## Tech Stack
- **Backend:** Golang (Fiber framework)
- **Database:** MySQL
- **Deployment:** Render (free-tier hosting)
- **Others:** Docker (optional), Railway (optional)

## Directory Structure
```
url-shortener/
│── auth-service/           # Authentication service
│   ├── database/           # DB connection & migrations
│   ├── middleware/         # Auth middleware
│   ├── models/             # User models
│   ├── utils/              # Helper functions
│   └── main.go             # Auth service entry point
│
│── url-service/            # URL shortening service
│   ├── database/           # DB connection & migrations
│   ├── models/             # URL models
│   ├── main.go             # URL service entry point
│
│── rate-limiting-service/  # Rate limiting service
│
│── cmd/                    # Main entry point
│   ├── main.go             # Starts all services
│
│── Dockerfile              # Containerization (optional)
│── go.mod                  # Go dependencies
│── go.sum                  # Dependency checksums
```

## Installation & Setup
### Prerequisites
- Go 1.20+
- MySQL database
- Git

### Clone the Repository
```sh
git clone https://github.com/EthicalMayor/url-shortener.git
cd url-shortener
```

### Install Dependencies
```sh
go mod tidy
```

### Environment Variables
Create a `.env` file with the following variables:
```env
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=url_shortener
PORT=5001
```

### Run Locally
```sh
go run cmd/main.go
```

## API Endpoints
### Shorten URL
```http
POST /shorten
```
#### Request Body
```json
{
  "original_url": "https://example.com"
}
```
#### Response
```json
{
  "short_url": "http://localhost:5001/abc123"
}
```

### Retrieve Original URL
```http
GET /{short_code}
```
#### Response
```json
{
  "original_url": "https://example.com"
}
```

## Deployment
### Deploy on Render
1. Push your code to GitHub.
2. Sign up at [Render](https://render.com).
3. Create a new **Web Service** and connect your GitHub repo.
4. Set the **Build Command:** `cd cmd && go build -o ../app`
5. Set the **Start Command:** `./app`
6. Add environment variables as per `.env`.
7. Deploy and get your public API URL!



