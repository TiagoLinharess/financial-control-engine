-- name: CreateTransaction :one
INSERT INTO transactions (
    user_id,
    name,
    date,
    value,
    category_id,
    credit_card_id,
    monthly_transactions_id,
    annual_transactions_id,
    installment_transactions_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE id = $1;

-- name: ListTransactionsByUserIDPaginated :many
SELECT 
    id,
    user_id,
    name,
    date,
    value,
    category_id,
    credit_card_id,
    monthly_transactions_id,
    annual_transactions_id,
    installment_transactions_id,
    created_at,
    updated_at,
    COUNT(*) OVER() as total_count
FROM transactions
WHERE user_id = $1
ORDER BY date DESC
LIMIT $2 OFFSET $3;

-- name: ListTransactionsByUserAndDate :many
SELECT 
    *,
    COUNT(*) OVER() as total_count
FROM transactions
WHERE user_id = $1
  AND date >= $2
  AND date <= $3
ORDER BY date DESC
LIMIT $4 OFFSET $5;

-- name: UpdateTransaction :one
UPDATE transactions
SET
    name = $2,
    date = $3,
    value = $4,
    category_id = $5,
    credit_card_id = $6,
    monthly_transactions_id = $7,
    annual_transactions_id = $8,
    installment_transactions_id = $9,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;