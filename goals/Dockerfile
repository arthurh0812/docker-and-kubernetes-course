FROM golang:1.15-alpine

WORKDIR /app

# the go.mod and go.sum package manager files
COPY go.mod go.mod
COPY app.go app.go

RUN export GO111MODULE=on

# copy the packages
COPY db/ db/

# install the dependencies
RUN go get

COPY . .

ENV DB_PASS=""
ENV DB_USER=""
ENV DB_NAME=""
ENV DB_HOST=0.0.0.0
ENV DB_PORT=5432

RUN CGO_ENABLED=0 go build -a -o app .

CMD [ "/app/app" ]