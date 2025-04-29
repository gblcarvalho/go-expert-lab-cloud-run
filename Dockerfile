# Etapa 1: build da aplicação
FROM golang:1.24 AS builder
WORKDIR /app

# Copia tudo para o build
COPY . .

# Compila o binário na pasta build
RUN go build -o /app/build/server /app/cmd/server/main.go

# Etapa 2: imagem final e enxuta
FROM debian:bookworm-slim
WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Cria diretório build (se necessário)
RUN mkdir -p /app/build

# Copia o .env (se existir)
COPY --from=builder /app/cmd/server/.env /app/.env

# Copia o binário
COPY --from=builder /app/build/server /app/build/server

# Expõe a porta do app
EXPOSE 8000

# Comando de execução
CMD ["./build/server"]
