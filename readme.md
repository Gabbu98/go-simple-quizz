# Aviation Quizz Instructions
## Author: Gabriel Camilleri 

A simple interactive quizz related to aviation. You will be presented with a series of multiple choice questions, where only one answer is expected.

## Instructions

- run: go run main.go, to display a list of cobra commands
- run: go run main.go start-api, to start the api on localhost:8080
- run: go run main.go start-quizz, to start the quizz

## Tech

Dillinger uses a number of open source projects to work properly:

- Golang
- Cobra CLi

The project is on the [public repository](https://github.com/Gabbu98/go-simple-quizz) on GitHub.

## Execution Order

Run the following executions in this order:

```sh
go run main.go start-api
```
```sh
go run main.go start-quizz
```