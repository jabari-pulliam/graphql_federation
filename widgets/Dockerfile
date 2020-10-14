FROM golang:1.14.2-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    GOPRIVATE=github.com/SemanticSugar/

WORKDIR /build

# Copy and download dependencies using go mod
RUN apk update && apk add git
COPY go.mod .
COPY go.sum .

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o app .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy dist files from build to server folder
RUN cp /build/app .

# Create test layer
FROM golang:1.14.2-alpine AS test

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPRIVATE=github.com/SemanticSugar/

COPY --from=builder /build /src

# Copy over the GOPATH to avoid re-downloading dependencies
COPY --from=builder /go /go

WORKDIR /src

# create development layer
FROM golang:1.14.2-alpine AS development

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# live reload
RUN apk update && apk add git

# Copy over the GOPATH to avoid re-downloading dependencies
COPY --from=builder /go /go

WORKDIR /src

EXPOSE 8080

# Create distribution layer
FROM alpine:3.11
COPY --from=builder /dist /

# Expose the port
EXPOSE 8080

CMD ["/app"]
