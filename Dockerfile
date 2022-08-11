FROM golang:1.18-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN go build

FROM golang:1.18-alpine
COPY --from=builder /build/go-pg-redis-memc /app/
WORKDIR /app
CMD ["./go-pg-redis-memc"]