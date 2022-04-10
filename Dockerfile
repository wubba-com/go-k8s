FROM golang:1.17-alpine as builder

WORKDIR /app
COPY go.mode .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /main.go

ENV PORT 8000

FROM alpine:3
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]