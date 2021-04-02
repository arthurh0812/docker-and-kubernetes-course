FROM golang:1.16.2 as builder

# Go environment variables
RUN GOCACHE=OFF
RUN go env -w GOPRIVATE=github.com/arthurh0812

WORKDIR /app

COPY go.mod go.sum app.go ./

ENV GOLANG_GITHUB_ACCESS_TOKEN=""

RUN git config --global url."https://golang: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@github.com".insteadOf "https://github.com"

RUN go get && go mod vendor
# creates a /app/vendor directory containing all the dependency code

COPY . .

ENV PORT=8080

RUN go build -o main


FROM scratch

COPY --from=builder /app /app

CMD [ "/app/main" ]