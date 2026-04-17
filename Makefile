MODULE   := github.com/gitlink-org/gitlink-cli
BINARY   := gitlink-cli
VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS  := -s -w -X '$(MODULE)/cmd.Version=$(VERSION)'

.PHONY: build install clean test

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) .

install:
	go install -ldflags "$(LDFLAGS)" .

clean:
	rm -f $(BINARY)

test:
	go test ./...
