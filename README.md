# Async File Reader

Async File Reader is a simple Go program that reads multiple files concurrently using goroutines and sync.WaitGroup. The program reads the contents of a given list of files asynchronously and returns their contents.

## Features

- Asynchronous file reading using goroutines.
- Uses sync.WaitGroup to wait for all goroutines to finish reading the files.
- Minimal and easy-to-understand code.

## Requirements

- Go 1.16 or later

## Usage

- Create your sample text files in the "files/" directory, or use the provided sample files.

- Run the application:
`go run ./cmd`
- To run the tests:
`go test ./...`
