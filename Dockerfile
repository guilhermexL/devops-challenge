# ==========================================
# Estágio 1: Build (Ambiente de compilação)
# ==========================================
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-server-projeto-looper .

# ==========================================
# Estágio 2: Final (Imagem de produção)
# ==========================================
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/http-server-projeto-looper .

EXPOSE 8080

CMD ["./http-server-projeto-looper"]
