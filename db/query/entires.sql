-- name: CreateEntry :one
INSERT INTO entires (
  account_id,
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entires 
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entires
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
UPDATE entires
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :one
DELETE FROM entires
WHERE id = $1
RETURNING *;

-- name: GetEntriesByAccount :many
SELECT * FROM entires
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

