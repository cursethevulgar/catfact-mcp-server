# catfact-mcp-server

MCP server to fetch a random cat fact from [catfact.ninja](https://catfact.ninja/).

## Overview

**catfact-mcp-server** is a lightweight Go-based server that exposes an endpoint for fetching random cat facts from the [catfact.ninja](https://catfact.ninja/) API. It is designed for easy integration with other services or as a fun standalone utility.

## Features

- Fetches random cat facts from an external API
- Simple and lightweight server, easy to deploy
- Written in Go for maximum performance

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher recommended)

### Installation

Clone the repository:

```bash
git clone https://github.com/cursethevulgar/catfact-mcp-server.git
cd catfact-mcp-server
```

### Usage

Build and run the server:

```bash
go run main.go
```

By default, the server listens on port **8080**.

#### Fetch a Cat Fact

Send a GET request to:

```
http://localhost:8080/catfact
```

You’ll receive a random cat fact as a JSON response.

### Example Response

```json
{
  "fact": "Cats have five toes on their front paws, but only four toes on their back paws."
}
```

## Configuration

You can modify the listening port or other parameters by editing `main.go`. (If environment variables or flags are supported, add those details here.)

## Project Structure

```
catfact-mcp-server/
├── main.go
└── README.md
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for improvements or new features.

## License

This project is licensed under the [MIT License](LICENSE).
