-- name: OrdersTask :one
INSERT INTO orders (
  trx_id,
  payment_trx_id,
  lock_items,
  status,
  customer_name,
  customer_address
)
VALUES (
  @trx_id,
  @payment_trx_id,
  @lock_items,
  @status,
  @customer_name,
  @customer_address
)
RETURNING id;