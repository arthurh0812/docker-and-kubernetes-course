FROM golang:1.16.3-alpine AS build

RUN apk add --no-cache git
RUN go env -w GO111MODULE=on

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod vendor && go mod verify

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/main .
# creates /tmp/build/main executable


FROM alpine:3.13

COPY --from=build /build/out/main /main

EXPOSE 80

CMD [ "/main" ]