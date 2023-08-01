### Home Broker

![Imersão Full Stack && Full Cycle](https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png)

Participe gratuitamente: https://imersao.fullcycle.com.br/

## Descrição do Projeto
 Imerção Full Cycle 13. Criando um Home Broker para comprar e vender ações.

## Tecnologias utilizadas
- Golang
- Docker
- Docker Compose
- Kafka
- MongoDB
- Prisma
- NestJS
- NextJS

## Como rodar o projeto

## Configure o docker host.docker.internal
- No Linux
```bash
$ sudo nano /etc/hosts
```
- No Windows
```bash
$ notepad C:\Windows\System32\drivers\etc\hosts
```
- Adicione a linha abaixo no final do arquivo
```
127.0.0.1 host.docker.internal
```
Use cat para verificar se a linha foi adicionada
```bash
$ cat /etc/hosts
```

### Clone o repositório
```bash
$ git clone https://github.com/FelipeCostaAraujo/imersao13.git
```

### Go
```bash
$ cd go
$ docker-compose up -d
$ go run cmd/trade/main.go
```
### NestJS
```bash
$ cd nestjs
$ docker-compose up -d
```

### NextJS
```bash
$ cd nextjs
$ docker-compose up -d
```

### Confluent
- Acesse o Confluent em http://localhost:9021
- Crie um tópico chamado `ORDERS_PUBLISHER`


## Fluxo de comunicação
```mermaid
sequenceDiagram
    participant User
    participant HomeBroker
    participant NestJS
    participant Kafka
    participant MicroservicesEmGo
    User->>HomeBroker: Negociação (React/Next.js)
    HomeBroker->>NestJS: Compra/venda/consultas (REST)
    NestJS->>Kafka: Publica ordens (compra/venda)
    Kafka->>MicroservicesEmGo: Consome ordens (compra/venda)
    MicroservicesEmGo->>Kafka: Publica matches das ordens
    Kafka->>NestJS: Consome ordens processadas
    NestJS->>HomeBroker: SSE (Service Sents Events)
    HomeBroker->>User: Recebe resposta da negociação 
```

## Acessando o projeto
- Home Broker: http://localhost:3001
- NestJS: http://localhost:3000
- Confluent: http://localhost:9021
- Prisma Studio: http://localhost:5555

## Acessando o banco de dados
- Mongo Compass: mongodb://root:root@127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.8.2&authSource=admin
