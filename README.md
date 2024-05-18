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
- [Supported languages](#supported-languages)
- [Declarative Comments Format](#declarative-comments-format)
  - [General Project Info](#general-project-info)
  - [General File Info](#general-file-info)
  - [General Function Info](#general-function-info)
  - [Backend explanation](#backend-explanation)

## Getting started

1. Add comments to your source code, See [Declarative Comments Format](#declarative-comments-format).

2. Install godoc by using:
```shell
go install github.com/mitsuaaki/app/cmd/app@latest 
```
To build from source you need [Go](https://golang.org/dl) (1.22 or newer).

Alternatively you can run the docker image:
```shell
make docker
```

Or download a pre-compiled binary from the [release page](https://github.com/mitsuaaki/godoc/releases).

3. Run `godoc init` in the root folder. This will parse your comments and generate the required files.
```shell
app init
```

## Supported languages

At the moment, godoc only supports Go.<br />
You can see some examples in the [example](https://github.com/mitsuaaki/godoc/example/) folder.

If you want to add support for another language, you can create a new issue or a pull request.

## Declarative Comments Format

### General Project Info

| Annotation           | Description                                         | Example                             |
|----------------------|-----------------------------------------------------|-------------------------------------|
| project-title        | **Required.** The title of the project.             | // @title Godoc Example             |
| project-version      | **Required.** Provides the version of the project.  | // @version 1.0                     |
| project-description  | A short description of the application.             | // @description This is a cool desc |

### General File Info

| Annotation       | Description                                        | Example                               |
|------------------|----------------------------------------------------|---------------------------------------|
| file-description | A short description of the file.                   | // @description This file contains... |

### General Function Info

| Annotation   | Description                          | Example                               |
|--------------|--------------------------------------|---------------------------------------|
| description  | A short description of the function. | // @description This function does... |
| param        | Description of a parameter.          | // @param a This is the first param   |
| return       | Description of the return value.     | // @return The result of the addition |
| complexity   | The complexity of the function.      | // @complexity O(n)                   |
| example      | Example code for using the function. | // @example `functionExample()`       |
| author       | Author of the function.              | // @author John Doe                   |

### Backend explanation

![image](https://github.com/mitsuaaki/godoc/assets/69150061/ac49811b-d87b-4cdf-b9d2-57e4f00ada88)

