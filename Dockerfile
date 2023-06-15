FROM golang:latest

WORKDIR /test

COPY . .

RUN go build -o /test

EXPOSE 8080

CMD ["/test"]