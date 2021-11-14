# Basic App Development

## Feature

- Account Service
- Order Service
    - Checkout
        - [] register to payment service
        - [] store and locking items
        - [] redirect to payment service page
    - Placed
        - [] update order to order placed
- Payment Service
    - Register Payment
        - store and locking items
    - Payment Page
        - choose payment type
        - simulating payment process
- Catalog Service

## Migrate

Run postgre on docker

```bash
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=root \
  -e POSTGRES_DB=commerce \
  -p 5432:5432 \
  --name commercedb \
  postgres:12.5-alpine
```

Export Variable

```bash
export POSTGRESQL_URL='postgres://root:root@localhost:5432/commerce?sslmode=disable'
```

### Create migrations

```bash
migrate create -ext sql -dir db/migrations/order -seq create_order_table
migrate create -ext sql -dir db/migrations/order -seq create_order_detail_table
```

### Run Migrations

```bash
migrate -database ${POSTGRESQL_URL} -path db/migrations/order up
```

## Tools

- Authentication & Authorization
- API Gateway

### Authentication & Authorization

### API Gateway

### OpenAPI

OpenAPI adalah Spesifikasi

Swagger adalah tools untuk mengimplementasi spesifikasi