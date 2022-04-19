FROM golang:1.17-alpine as build

WORKDIR /app
COPY go.mod go.mod
#COPY go.sum go.sum
RUN go mod download
COPY ./ ./

RUN go build -o main main.go

CMD ["./main"]