# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /src
RUN apk add --no-cache ca-certificates git

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build (static)
ARG GIT_SHA=dev
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-s -w -X 'github.com/dangrondahl/hello-kosli/internal/version.GitSHA=${GIT_SHA}'" \
    -o /out/hello-kosli ./cmd/hello-kosli

# Minimal image
FROM gcr.io/distroless/static-debian12
ENV PORT=8080
EXPOSE 8080
COPY --from=build /out/hello-kosli /hello-kosli
USER nonroot:nonroot
ENTRYPOINT ["/hello-kosli"]
