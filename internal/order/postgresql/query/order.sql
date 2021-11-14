-- name: CreateOrders :one
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


-- name: OrderPlaced :one
UPDATE orders SET
  status = @status
WHERE payment_trx_id = @payment_trx_id
RETURNING id AS res;

-- name: SelectPayloads :one
SELECT id, lock_items 
FROM orders
WHERE payment_trx_id = @payment_trx_id;

-- name: CreateOrdersDetail :one
INSERT INTO order_detail (
  order_id,
  name,
  quantity,
  price
)
VALUES (
  @order_id,
  @name,
  @quantity,
  @price
)
RETURNING id;