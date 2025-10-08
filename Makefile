export GOEXPERIMENT=jsonv2

test:
	go test ./...

build:
	go build .

run:
	go run main.go
