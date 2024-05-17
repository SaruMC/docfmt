<p><a target="_blank" href="https://app.eraser.io/workspace/a5WYfKVeq8Sek7ssM9Fc" id="edit-in-eraser-github-link"><img alt="Edit in Eraser" src="https://firebasestorage.googleapis.com/v0/b/second-petal-295822.appspot.com/o/images%2Fgithub%2FOpen%20in%20Eraser.svg?alt=media&amp;token=968381c8-a7e7-472a-8ed6-4a6626da5501"></a></p>

![logo](https://github.com/mitsuaaki/godoc/assets/69150061/25510cf3-17ca-44d2-93bb-81b698eb4504 "")

# godoc
Godoc generates a formatted PDF of your code's documentation also name docstring.
It allows you to have a documentation that's easy to read / access for every developer that can work on your project.

## Context
When I was working on an API rest for a personal project, I came up against the problem of documentation.
Well-known solutions like [﻿swagger](https://swagger.io/) exist.
Nowadays, however, public documentation can be a source of vulnerabilities.
To protect my infrastructure from fuzzing, I set myself the challenge of creating a PDF documentation generated from the docstring of any project.

This is how godoc was born.
And this is also why for the moment, this README really looks like the one from a cool project named [﻿Swaggo](https://github.com/swaggo).

## Contents
- [﻿Context](#context) 
- [﻿Getting started](#getting-started) 
- [﻿Declarative Comments Format](#declarative-comments-format) 
## Getting started
1. Add comments to your source code, See [﻿Declarative Comments Format](#declarative-comments-format) .
2. Install godoc by using:
```shell
go install github.com/mitsuaaki/godoc/cmd/godoc@latest
```
To build from source you need [﻿Go](https://golang.org/dl) (1.22 or newer).

Alternatively you can run the docker image:

```shell
make docker
```
Or download a pre-compiled binary from the [﻿release page](https://github.com/mitsuaaki/godoc/releases).

1. Run `godoc init`  in the root folder. This will parse your comments and generate the required files.
```shell
godoc init
```
## Declarative Comments Format
### General Project Info
| Annotation | Description | Example |
| ----- | ----- | ----- |
| title | <p>**Required.**</p><p> The title of the project.</p> | // @title Godoc Example |
| version | <p>**Required.**</p><p> Provides the version of the project.</p> | // @version 1.0 |
| description | A short description of the application. | // @description This is a cool desc |
### General File Info
| Annotation | Description | Example |
| ----- | ----- | ----- |
| description | A short description of the file. | // @description This file contains... |
### General Function Info
| Annotation | Description | Example |
| ----- | ----- | ----- |
| description | A short description of the function. | // @description This function does... |
| param | Description of a parameter. | // @param a This is the first param |
| return | Description of the return value. | // @return The result of the addition |
| complexity | The complexity of the function. | // @complexity O(n) |
| example | Example code for using the function. | // @example  |
| author | Author of the function. | // @author John Doe |
### Backend explanation
![Backend](/.eraser/a5WYfKVeq8Sek7ssM9Fc___Cs3WexGEn0O5x35bLyqRQh5HCKC3___---figure---qAMSi0K-iPxm5n6Pn_Lrn---figure---00QXBfgA52UabrJ0HB9VHg.png "Backend")




<!-- eraser-additional-content -->
## Diagrams
<!-- eraser-additional-files -->
<a href="/README-cloud-architecture-1.eraserdiagram" data-element-id="xfWs_uQ-nnkJIhcClZFp6"><img src="/.eraser/a5WYfKVeq8Sek7ssM9Fc___Cs3WexGEn0O5x35bLyqRQh5HCKC3___---diagram----de5b4238cf653b20154ba1d0652501c4.png" alt="" data-element-id="xfWs_uQ-nnkJIhcClZFp6" /></a>
<!-- end-eraser-additional-files -->
<!-- end-eraser-additional-content -->
<!--- Eraser file: https://app.eraser.io/workspace/a5WYfKVeq8Sek7ssM9Fc --->