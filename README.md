# go-expert-lab-cloud-run


Implementação de um desafio de Cloud Run.
🛠️ Pré-requisitos

Antes de executar o projeto, certifique-se de ter os seguintes softwares instalados:

- Docker
- Docker Compose

## 🚀 Como rodar o projeto
1. Subindo o ambiente com Docker Compose

Execute os seguintes comando apra subir o ambiente via docker compose.

Primeiro é necessário criar um arquivo .env contendo as variáveis de ambiente. Ele pode ser criado copiando o .env.exemple e alterando o valor de WEATHER_API_KEY

```bash
cp ./cmd/server/.env.exemple ./cmd/server/.env
```

Depois é só executar o comando abaixo passando o arquivo .env criado como parâmetro

```bash
docker compose --env-file ./cmd/server/.env up -d
```

Este comando irá:

- Iniciar o servidor da applicação Go


O serviço da aplicação está disponíveis em
- HTTP (RESTful): http://localhost:8000/weather/{cep}


## 🧪 Testes automatizados

Para facilitar, o projeto também fornece um comando para rodar testes automáticos que disparam múltiplas requisições com CEPs válidos e inválidos:

```bash
source ./test.sh
```
