# Basic App Development

Ecommerce workflow sederhana, handle order dan payment page.

## Feature

- Order Service
    - Checkout
        - [x] register to payment service
        - [x] store and locking items
        - [x] redirect to payment service page
    - Placed
        - [x] update order to order placed
- Payment Service
    - Register Payment
        - [x] store and locking items
    - Payment Page
        - [x] choose payment type
        - [x] simulating payment process
- Account Service
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
