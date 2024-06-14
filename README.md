<img align="right" width="300px" src="https://github.com/mitsuaaki/docfmt/assets/69150061/25510cf3-17ca-44d2-93bb-81b698eb4504" alt="logo">

# docs-formatter

[![Build Status](https://github.com/mitsuaaki/docfmt/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/features/actions)
[![codebeat badge](https://codebeat.co/badges/c8c55d12-43ca-40dd-9901-cf87a472273b)](https://codebeat.co/projects/github-com-mitsuaaki-docfmt-main)
[![Release](https://img.shields.io/github/release/mitsuaaki/docfmt.svg?style=flat-square)](https://github.com/mitsuaaki/docfmt/releases)


docs-formatter generates a formatted PDF of your code's documentation also name docstring.
It allows you to have a documentation that's easy to read / access for every developer that can work on your project.

## Context

When I was working on an API rest for a personal project, I came up against the problem of documentation. 
Well-known solutions like [swagger](https://swagger.io/) exist. 
Nowadays, however, public documentation can be a source of vulnerabilities. 
To protect my infrastructure from fuzzing, I set myself the challenge of creating a PDF documentation generated from the docstring of any project.

This is how docs-formatter was born.
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

2. Install docs-formatter by using:
```shell
go install github.com/mitsuaaki/docs-formatter/cmd/docs-formatter@latest 
```
To build from source you need [Go](https://golang.org/dl) (1.22 or newer).

Or download a pre-compiled binary from the [release page](https://github.com/mitsuaaki/docs-formatter/releases).

3. Run `docs-formatter init` in the root folder. This will parse your comments and generate the required files.
```shell
fdocs init
```

For the moment, docs-formatter only supports a few languages. <br>
If you want to add support for another language, you can create a new issue or a pull request.<br>

Also, only testing purposes command or available:
```shell
fdocs test-generate <options>
fdocs test-scan <options>
```

If you want to contribute to the PDF part, that would be great. <br>
You can find the code in the following folder: [pdf](https://github.com/mitsuaaki/docs-formatter/app/formatter). 

(Help me, it's really pain in the ass however it's not that hard just boring)

## Supported languages

At the moment, docs-formatter only supports Go.<br />
You can see some examples in the [example](https://github.com/mitsuaaki/docs-formatter/example/) folder.

If you want to add support for another language, you can create a new issue or a pull request.

## Declarative Comments Format

### General Project Info

| Annotation           | Description                                         | Example                             |
|----------------------|-----------------------------------------------------|-------------------------------------|
| project-title        | **Required.** The title of the project.             | // @title docs-formatter Example             |
| project-version      | **Required.** Provides the version of the project.  | // @version 1.0                     |
| project-description  | A short description of the application.             | // @description This is a cool desc |

### General File Info

| Annotation       | Description                                        | Example                               |
|------------------|----------------------------------------------------|---------------------------------------|
| file-description | A short description of the file.                   | // @description This file contains... |

### General Function Info

| Annotation   | Description                          | Example                                      |
|--------------|--------------------------------------|----------------------------------------------|
| description  | A short description of the function. | // @description This function does...        |
| param        | Description of a parameter.          | // @param a - type - This is the first param |
| return       | Description of the return value.     | // @return The result of the addition        |
| complexity   | The complexity of the function.      | // @complexity O(n)                          |
| example      | Example code for using the function. | // @example `functionExample()`              |
| author       | Author of the function.              | // @author John Doe                          |

### Backend explanation

![image](https://github.com/mitsuaaki/docs-formatter/assets/69150061/ac49811b-d87b-4cdf-b9d2-57e4f00ada88)

