# Microservices com gRPC - Parte 1

Este projeto demonstra a implementação de dois microsserviços utilizando gRPC em Go: um serviço de pedidos (orders) e um serviço de pagamentos (payments), ambos com persistência em MySQL. O serviço de pedidos se comunica com o serviço de pagamentos via gRPC para processar pagamentos de pedidos.

## Pré-requisitos

- Go 1.19 ou superior
- Docker e Docker Compose
- Git

## Instalação e Configuração

### 1. Clone o repositório

```bash
git clone <url-do-repositorio>
cd pratica-microservicos-com-grpc-parte-1
```

### 2. Instale as dependências do Go

```bash
cd microservices/order
go mod tidy
```

### 3. Gere o código protobuf (opcional, se já não estiver gerado)

Execute o script `run.sh` na raiz do projeto para gerar o código protobuf do serviço de pedidos:

```bash
./run.sh
```

Este script irá:
- Instalar o protoc-gen-go e protoc-gen-go-grpc
- Gerar o código Go a partir do arquivo `order.proto`

Para o serviço de pagamentos, execute o script similarmente (se necessário):
```bash
SERVICE_NAME=payment ./run.sh
```

## Como Executar

### Opção 1: Usando Docker Compose (Recomendado)

Esta é a maneira mais fácil de executar todo o sistema com todos os serviços.

#### 1. Execute todos os serviços

```bash
docker-compose up --build
```

Este comando irá:
- Construir e iniciar o container MySQL
- Construir e iniciar o serviço de pagamentos (porta 3001)
- Construir e iniciar o serviço de pedidos (porta 3000)

#### 2. Teste o sistema

Em outro terminal, execute o cliente de teste:

```bash
cd microservices/order
go run client/main.go
```

O cliente irá:
- Conectar ao servidor gRPC de pedidos na porta 3000
- Criar um pedido de exemplo
- O serviço de pedidos irá se comunicar com o serviço de pagamentos para processar o pagamento
- Exibir o ID do pedido criado

#### 3. Parar os serviços

```bash
docker-compose down
```

### Opção 2: Execução Manual (Desenvolvimento)

Se preferir executar manualmente para desenvolvimento:

#### 1. Inicie o banco de dados MySQL

```bash
docker-compose up mysql -d
```

#### 2. Execute o serviço de pagamentos

```bash
cd microservices/payment
go run cmd/main.go
```

#### 3. Execute o serviço de pedidos (em outro terminal)

```bash
cd microservices/order
go run cmd/main.go
```

#### 4. Execute o cliente de teste (em outro terminal)

```bash
cd microservices/order
go run client/main.go
```

## Arquitetura

O projeto consiste em dois microsserviços:

- **Serviço de Pedidos (Order Service)**: Gerencia a criação e processamento de pedidos. Porta: 3000
- **Serviço de Pagamentos (Payment Service)**: Processa pagamentos para os pedidos. Porta: 3001

O serviço de pedidos se comunica com o serviço de pagamentos via gRPC para validar e processar pagamentos.

## Variáveis de Ambiente

### Serviço de Pedidos
- `DATA_SOURCE_URL`: URL de conexão com o banco MySQL para o banco `order_db`
  - Formato: `user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local`
- `APPLICATION_PORT`: Porta onde o servidor gRPC irá escutar (padrão: 3000)
- `PAYMENT_SERVICE_URL`: URL do serviço de pagamentos (ex: `payment:3001`)
- `ENV`: Ambiente de execução (development/production)

### Serviço de Pagamentos
- `DATA_SOURCE_URL`: URL de conexão com o banco MySQL para o banco `payment`
  - Formato: `user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local`
- `APPLICATION_PORT`: Porta onde o servidor gRPC irá escutar (padrão: 3001)
- `ENV`: Ambiente de execução (development/production)

