FROM golang:1.22 AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app

# end of build stage

FROM alpine:latest
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder app .

EXPOSE 8080
EXPOSE 8090

ENTRYPOINT ["./app"]