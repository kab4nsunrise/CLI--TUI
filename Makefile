.PHONY: build run test tidy clean install

BINARY=gomon
CMD=./cmd/gomon

build:
	go build -o bin/$(BINARY) $(CMD)

run: build
	./bin/$(BINARY)

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/

install: build
	cp bin/$(BINARY) $(GOPATH)/bin/$(BINARY)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY)-linux-amd64 $(CMD)

build-darwin:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY)-darwin-arm64 $(CMD)