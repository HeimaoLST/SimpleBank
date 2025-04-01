-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers 
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET from_account_id = $2,
    to_account_id = $3,
    amount = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :one
DELETE FROM transfers
WHERE id = $1
RETURNING *;

-- name: GetTransfersByFromAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id;

-- name: GetTransfersByToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id;

-- name: GetTransfersByFromAndToAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1 AND to_account_id = $2
ORDER BY id;


