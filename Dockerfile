FROM golang:1.22.1

WORKDIR /
COPY go.mod .
COPY go.sum .

COPY / .

RUN go mod tidy
RUN go build -o main ./cmd/main.go


CMD ["main"]
