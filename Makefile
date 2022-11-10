PROGRAM=autowar

build: clean test static
	go build -o $(PROGRAM) cmd/main.go

test:
	go test  ./...

static:
	staticcheck ./...

clean:
	rm -f $(PROGRAM) results.csv