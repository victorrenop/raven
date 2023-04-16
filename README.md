<img align="right" width="150" height="150" src="assets/raven-logo.png">

# Raven Config Manager
Raven is a configuration management tool written in Golang that provides a way to manage app configurations through an easy API. It is designed to be easy to use and highly configurable and it's mostly a fun project written in Golang for internal uses in our Discord Bot.

## Getting Started
To get started with Raven, simply clone this repository and install the dependencies using your preferred package manager. Raven is built using Golang.

```bash
Copy code
git clone https://github.com/your-username/raven-api.git
cd raven-api
go mod download
```
Once you have installed the dependencies, you can start the Raven API server using the following command:

```go
Copy code
go run main.go
```

Raven comes with a docker configuration too! Simply build the docker image and run it using your preferred options.

## API Documentation
Most of Raven's documentation is located in the swagger file using the OpenAPI 3.0 specification.

## License
Raven is released under the MIT License. Feel free to use and modify the code as you see fit.