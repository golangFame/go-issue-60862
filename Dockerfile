FROM golang:1.20.5-alpine3.18

WORKDIR /go/src/github.com/riandyrn/try-otelchi

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build -o ./app

ENTRYPOINT [ "./app" ]