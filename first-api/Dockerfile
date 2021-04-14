FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .

RUN go mod vendor

COPY . .

EXPOSE 443/tcp

RUN go build . -o main

CMD [ "main" ]