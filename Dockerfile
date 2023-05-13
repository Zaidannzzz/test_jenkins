FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o /test/backend

EXPOSE 8080

CMD ["/"]