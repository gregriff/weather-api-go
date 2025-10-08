# weather-api-go

This project is a rewrite of the weather-api repo written in Python. I'm doing this to learn how to develop HTTP APIs in Go and continue to learn the language. This repo is being developed with the intention of being a good skeleton for future HTTP API projects.

### Objective:
Feature parity with [the python weather API](https://github.com/gregriff/weather-api), but written using only the Go standard library.

### Features:

##### Validation:
- HTTP request and response bodies are validated using a very simple system detailed [here](./internal/validation/Validation.md)
