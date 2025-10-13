APP           := hello-kosli
MODULE        := github.com/dangrondahl/hello-kosli
GIT_SHA       := $(shell git rev-parse --short HEAD 2>/dev/null || echo dev)
LDFLAGS       := -s -w -X '$(MODULE)/internal/version.GitSHA=$(GIT_SHA)'
GO            := go

.PHONY: all build test run tidy clean docker-build docker-run

all: build

build:
	$(GO) build -ldflags "$(LDFLAGS)" -o bin/$(APP) ./cmd/hello-kosli

test:
	$(GO) test ./... -cover

run:
	PORT=8080 $(GO) run -ldflags "$(LDFLAGS)" ./cmd/hello-kosli

tidy:
	$(GO) mod tidy

clean:
	rm -rf bin

docker-build:
	docker build --build-arg GIT_SHA=$(GIT_SHA) -t $(APP):$(GIT_SHA) .

docker-run:
	docker run --rm -p 8080:8080 -e PORT=8080 $(APP):$(GIT_SHA)
