Simple HTTP Server in Go

This is a simple HTTP server written in Go that supports basic GET and POST requests for handling files and echoing messages.

Features:
Handles GET requests for:

Root ("/")

Echoing messages ("/echo/{message}")

User agent ("/user-agent")

Serving files ("/files/{filename}")

Supports POST requests for:

Uploading files ("/files/{filename}")

Make requests:

GET Request Examples:

curl http://localhost:4221/

curl http://localhost:4221/echo/hello

curl http://localhost:4221/user-agent

curl http://localhost:4221/files/{filename}

POST Request Example:

curl -X POST -d "data_to_upload" http://localhost:4221/files/{filename}

Credits to https://app.codecrafters.io/courses/http-server

Prerequisites:
Go 1.16 or higher installed
Operating system: Linux, macOS, or Windows
