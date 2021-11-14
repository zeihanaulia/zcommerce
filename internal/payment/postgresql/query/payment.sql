-- name: RegisterPayments :one
INSERT INTO payments (
  trx_id,
  reference_trx_id,
  types,
  final_amount,
  payloads,
  status
)
VALUES (
  @trx_id,
  @reference_trx_id,
  @types,
  @final_amount,
  @payload,
  @status
)
RETURNING id;


-- name: SelectPayments :one
SELECT
  trx_id,
  reference_trx_id,
  types,
  final_amount,
  payloads,
  status
FROM
  payments
WHERE
  trx_id = @trx_id
LIMIT 1;

-- name: PaidPayment :one
UPDATE payments SET
  types = @types,
  status = @status
WHERE trx_id = @trx_id
RETURNING id AS res;