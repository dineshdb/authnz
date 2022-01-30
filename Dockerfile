FROM docker.io/golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go build -o server cmd/server/main.go

FROM docker.io/ubuntu
WORKDIR /app
COPY --from=builder /app/server /usr/local/bin/server
CMD ["server"]