# go Battleship 

A simple game of Battleship, written in golang.

# Getting started

This project requires a go v1.16 or higher. To prepare to work with it, pick one of these options:

## Run locally

Run battleship 

```bash
go run cmd/main.go
```

Execute tests 

```bash
go test ./...
```

## Docker

If you don't want to install anything related with golang on your system, you can
run the game inside Docker instead.

### Run a Docker Container from the Image

```bash
docker run -it -v ${PWD}:/battleship -w /battleship golang:1.16 bash
```

This will run a Docker container with your battleship case study mounted into it. The container will run in interactive mode and you can execute Gradle commands from the shell (see examples below).

# Launching the game

```bash
go run main.go
```

# Running the Tests

```bash
go test
```
