FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build

FROM alpine

COPY --from=builder /build/app-service /app-service

EXPOSE 8080

ENTRYPOINT ["/app-service"]
