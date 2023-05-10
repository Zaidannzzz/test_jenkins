FROM golang:latest

WORKDIR /app

RUN go mod download
COPY . .

RUN go build -o test-api

EXPOSE 8080

CMD ["./main"]