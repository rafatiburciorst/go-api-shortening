# Shortening

Shortening is a URL shortening service written in Go. It provides a simple API to shorten URLs and redirect to the original URLs.

## Project Structure
api/ api.go go.mod go.sum main.go readme.md


## Getting Started

### Prerequisites

- Go 1.22.6 or later

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/encurtador.git
    cd encurtador
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

### Running the Server

To start the server, run:

```sh
go run [main.go](http://_vscodecontentref_/#%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Frafaeltiburcio%2Fdev%2Flanguages%2Fgo%2Fapis%2Fprojeto-api%2Fmain.go%22%2C%22path%22%3A%22%2FUsers%2Frafaeltiburcio%2Fdev%2Flanguages%2Fgo%2Fapis%2Fprojeto-api%2Fmain.go%22%2C%22scheme%22%3A%22file%22%7D%7D)

```
The server will start on localhost:8080.

### API Endpoints
#### Shorten URL
URL: /api/shorten
Method: POST
Request Body:
```
{
    "url": "https://example.com"
}
```
#### Response
``` 
{
    "data": "shortenedCode"
}
```

### Redirect to Original URL
- URL: /{code}
- Method: GET
### Code Overview
#### Main Entry Point
The main entry point is in main.go. The run function sets up the HTTP server and routes.

#### API Handlers
API handlers are defined in api/api.go.

- NewHandler: Sets up the routes and middleware.
- handlePost: Handles URL shortening requests.
- handleGet: Handles redirection to the original URL.
- sendJSON: Sends JSON responses.
- genCode: Generates a random code for shortened URLs.
### Dependencies
go-chi/chi: Lightweight, idiomatic and composable router for building Go HTTP services.
### License
This project is licensed under the MIT License.