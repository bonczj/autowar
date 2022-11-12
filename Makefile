PROGRAM=autowar

build: clean test static
	go build -o $(PROGRAM) cmd/main.go

test:
	go test  ./...

static: 
	# if the `staticcheck` binary does not exist, install it
	which staticcheck > /dev/null || \
		(go install honnef.co/go/tools/cmd/staticcheck@latest && \
		echo "Make sure to add staticcheck to your path. It is probably in ~/go/bin/")

	staticcheck ./...

clean:
	rm -f $(PROGRAM) results.csv