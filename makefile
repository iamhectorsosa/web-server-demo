all: build

run:
	go run .

test:
	go test -count=1 -v ./...
