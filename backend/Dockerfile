FROM golang:latest

WORKDIR /test/backend

COPY . .

ENV GO111MODULE=on

RUN go build -o /main

EXPOSE 8080

CMD ["./main"]