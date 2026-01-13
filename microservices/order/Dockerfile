FROM golang:1.24 AS builder
WORKDIR /usr/src/app
COPY . .
WORKDIR /usr/src/app/microservices/order
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o order ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/microservices/order/order ./order
CMD ["./order"]
