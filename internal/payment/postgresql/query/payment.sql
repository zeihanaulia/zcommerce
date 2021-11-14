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