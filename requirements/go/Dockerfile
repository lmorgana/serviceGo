FROM golang:1.18

WORKDIR /app

EXPOSE 8080

COPY ./*.go .
COPY ./go.mod ./

RUN go mod tidy

CMD go run *.go