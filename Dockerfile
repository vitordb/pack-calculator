FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o pack-calculator ./cmd/main.go

FROM golang:1.24-alpine

WORKDIR /app

COPY --from=builder /app/pack-calculator .

COPY --from=builder /app/ui ./ui

EXPOSE 8080

CMD ["./pack-calculator"]