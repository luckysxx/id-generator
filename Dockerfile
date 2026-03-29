# syntax=docker/dockerfile:1
FROM golang:1.25.5-alpine AS builder
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o id-generator cmd/idgen/main.go

FROM scratch
COPY --from=builder /app/id-generator /id-generator
EXPOSE 50059
ENTRYPOINT ["/id-generator", "-port=50059", "-node=1"]
