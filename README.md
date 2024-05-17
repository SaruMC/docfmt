<img align="right" width="300px" src="https://github.com/mitsuaaki/godoc/assets/69150061/25510cf3-17ca-44d2-93bb-81b698eb4504" alt="logo">

# godoc

[![Build Status]()]()

Godoc generates a formatted PDF of your code's documentation also name docstring.
It allows you to have a documentation that's easy to read / access for every developer that can work on your project.

## Context

When I was working on an API rest for a personal project, I came up against the problem of documentation. 
Well-known solutions like [swagger](https://swagger.io/) exist. 
Nowadays, however, public documentation can be a source of vulnerabilities. 
To protect my infrastructure from fuzzing, I set myself the challenge of creating a PDF documentation generated from the docstring of any project.

This is how godoc was born.
And this is also why for the moment, this README really looks like the one from a cool project named [Swaggo](https://github.com/swaggo).

## Contents
- [Context](#context)
- [Getting started](#getting-started)
- [Declarative Comments Format](#declarative-comments-format)

## Getting started

1. Add comments to your source code, See [Declarative Comments Format](#declarative-comments-format).

2. Install godoc by using:
```shell
go install github.com/mitsuaaki/godoc/cmd/godoc@latest 
```
To build from source you need [Go](https://golang.org/dl) (1.22 or newer).

Alternatively you can run the docker image:
```shell
make docker
```

Or download a pre-compiled binary from the [release page](https://github.com/mitsuaaki/godoc/releases).

3. Run `godoc init` in the root folder. This will parse your comments and generate the required files.
```shell
godoc init
```

## Declarative Comments Format

### General Project Info

| Annotation            | Description                                        | Example                             |
|-----------------------|----------------------------------------------------|-------------------------------------|
| title                 | **Required.** The title of the project.            | // @title Godoc Example             |
| version               | **Required.** Provides the version of the project. | // @version 1.0                     |
| description           | A short description of the application.            | // @description This is a cool desc |


### General File Info

| Annotation            | Description                                        | Example                             |
|-----------------------|----------------------------------------------------|-------------------------------------|
| description           | A short description of the application.            | // @description This is a cool desc |


## General Function Info

| Annotation  | Description                             | Example                             |
|-------------|-----------------------------------------|-------------------------------------|
| description | A short description of the application. | // @description This is a cool desc |


### Backend explanation

[!]()
