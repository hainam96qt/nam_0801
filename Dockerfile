FROM golang:1.21.4 AS go_builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o run

FROM alpine:3.16

WORKDIR /app

COPY --from=go_builder ./app/run ./nam-0801

RUN chmod +x /app/nam-0801

CMD ["/app/nam-0801"]
