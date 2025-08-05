FROM golang:1.22-alpine AS builder

RUN apk add --no-cache protobuf-dev git
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY proto/ ./proto/
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/vusb.proto

COPY main.go ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o hello-world .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/hello-world .
EXPOSE 50051
CMD ["./hello-world"]
