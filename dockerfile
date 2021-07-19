FROM golang:alpine AS builder

WORKDIR /app

ENV GOOS=linux
ENV GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o ./authentication/authsvc authentication/main.go
RUN GOOS=linux GOARCH=amd64 go build -o ./api/apisvc api/main.go

FROM golang:alpine

WORKDIR /app

COPY --from=builder /app/authentication/authsvc /app/api/apisvc /app/

EXPOSE 9000
EXPOSE 9001

# CMD ["/app/authsvc", "/app/apisvc"]