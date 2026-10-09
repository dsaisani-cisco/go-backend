FROM golang:1.24.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY cmd/ ./cmd/
COPY pkg/ ./pkg//

RUN go build -o main cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]