# URL Shortener

A simple URL shortener application built with Go. This application allows you to shorten long URLs and redirect to the original URLs using short codes.

## Features

- Shorten long URLs and receive a short URL
- Redirect to the original URL when accessing the short URL
- Basic health check endpoint to verify server status

## Getting Started

### Prerequisites

- Go (1.18 or higher) installed on your machine

### Installation

1. **Clone the Repository:**

   ```sh
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener

2. **Run the Application:**

```sh
    go run main.go
```

The server will start on http://localhost:8080.

3. **API Endpoints:**
 **Shorten URL** 
- **URL:** `http://localhost:8080/shorten`
- **Method:** POST
- **Request Body:**
  ```json
  {
    "url": "https://www.example.com"
  }





