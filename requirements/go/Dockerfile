FROM golang:1.18-alpine AS build

WORKDIR /src/
COPY *.go /src/
COPY ./go.mod ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
EXPOSE 8080

ENTRYPOINT ["/bin/demo"]