# go Battleship 

A simple game of Battleship, written in golang. The purpose of this repository is to serve as an entry point into coding exercises and it was especially created for scrum.orgs Applying Professional Scrum for Software Development course (www.scrum.org/apssd). The code in this repository is unfinished by design.

# Getting started

This project requires a go v1.17 or higher. To prepare to work with it, pick one of these options:

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
docker run -it -v ${PWD}:/battleship -w /battleship golang:1.17 bash
```

This will run a Docker container with your battleship case study mounted into it. The container will run in interactive mode and you can execute Gradle commands from the shell (see examples below).

# Launching the game

```bash
go run cmd/main.go
```

# Running the Tests

```bash
go test ./...
```
