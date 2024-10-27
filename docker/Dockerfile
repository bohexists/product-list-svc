FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../ .

RUN go build -o product-list-svc ./cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/product-list-svc .

CMD ["./product-list-svc"]