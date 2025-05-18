FROM golang:1.24.3-alpine AS build
ARG VERSION="dev"

WORKDIR /build

RUN --mount=type=cache,target=/var/cache/apk \
    apk add git

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=${VERSION} -X main.commit=$(git rev-parse HEAD) -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -o /bin/catfact-mcp-server main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /server
COPY --from=build /bin/catfact-mcp-server .
CMD ["./catfact-mcp-server", "stdio"]
