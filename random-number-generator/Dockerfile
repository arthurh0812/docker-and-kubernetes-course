FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod .

COPY rng.go .

RUN go get

COPY . .

RUN go build -o rng

CMD ["/app/rng"]