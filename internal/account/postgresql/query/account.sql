-- name: CreateAccounts :one
INSERT INTO accounts (
  name,
  email,
  password
)
VALUES (
  @name,
  @email,
  @password
)
RETURNING id;

-- name: SelectAccounts :one
SELECT id, name, email, password 
FROM accounts
WHERE email = @email;
