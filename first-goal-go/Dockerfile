FROM golang:1.16-alpine

WORKDIR /app
# inside /app
COPY go.mod .

COPY server.go .

RUN go get

COPY . .

RUN go build -o server

EXPOSE 80

# "server" binary
CMD ["/app/server"]