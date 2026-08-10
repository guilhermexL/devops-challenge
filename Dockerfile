# ==========================================
# Estágio 1: Build (Ambiente de compilação)
# ==========================================
FROM golang:alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o restante do código fonte
COPY . .

# Compila a aplicação Go. 
# CGO_ENABLED=0 gera um binário estático, ideal para containers Alpine/Scratch.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-server-projeto-korp .

# ==========================================
# Estágio 2: Final (Imagem de produção)
# ==========================================
FROM alpine:latest

# Instala certificados de segurança caso a API precise fazer chamadas HTTPS no futuro
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia apenas o binário compilado do estágio anterior
COPY --from=builder /app/http-server-projeto-korp .

# Expõe a porta que o serviço utiliza
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./http-server-projeto-korp"]
