FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o test-api

EXPOSE 8080

CMD ./test-api