FROM golang:alpine

WORKDIR /test/backend

COPY . .

RUN go mod tidy

ENV GO111MODULE=on

RUN go build -o main 

EXPOSE 8080

CMD ["./main"]