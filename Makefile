APP           := hello-kosli
MODULE        := github.com/dangrondahl/hello-kosli
GIT_SHA       := $(shell git rev-parse --short HEAD 2>/dev/null || echo dev)
LDFLAGS       := -s -w -X '$(MODULE)/internal/version.GitSHA=$(GIT_SHA)'
GO            := go


COVER_DIR     := coverage
COVER_OUT     := $(COVER_DIR)/cover.out
COVER_JSON    := $(COVER_DIR)/coverage.json


.PHONY: all build test run tidy clean docker-build docker-run coverage verify-coverage

all: build

build:
	$(GO) build -ldflags "$(LDFLAGS)" -o bin/$(APP) ./cmd/hello-kosli

test:
	@mkdir -p $(COVER_DIR)
	$(GO) test ./... -covermode=atomic -coverpkg=./... -coverprofile=$(COVER_OUT)

coverage: test
	@total=$$(go tool cover -func=$(COVER_OUT) | grep '^total:' | awk '{print $$3}' | sed 's/%//'); \
	echo "{ \"coverage\": $$total }" | tee $(COVER_JSON)

run:
	PORT=8080 $(GO) run -ldflags "$(LDFLAGS)" ./cmd/hello-kosli

tidy:
	$(GO) mod tidy

clean:
	rm -rf bin $(COVER_DIR)

docker-build:
	docker build --build-arg GIT_SHA=$(GIT_SHA) -t $(APP):$(GIT_SHA) .

docker-run:
	docker run --rm -p 8080:8080 -e PORT=8080 $(APP):$(GIT_SHA)
