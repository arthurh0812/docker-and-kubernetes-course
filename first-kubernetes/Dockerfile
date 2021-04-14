FROM golang:1.16.3-alpine as build

# Go environment variables
RUN GOCACHE=OFF
RUN go env -w GOFLAGS=-mod=vendor
RUN go env -w GO111MODULE=on
RUN go env -w GOPRIVATE=github.com/arthurh0812

WORKDIR /app

COPY go.mod go.sum app.go ./

RUN go mod download && go mod vendor && go mod verify
# creates a /app/vendor directory containing all the dependency code

COPY . .

ENV PORT=8080

RUN go build -o main


FROM debian:buster

COPY --from=build /app /app

CMD [ "/app/main" ]