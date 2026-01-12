-- name: CreateMonthlyTransaction :one
INSERT INTO monthly_transactions (
    user_id,
    name,
    value,
    day,
    category_id,
    credit_card_id
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetMonthlyTransactionByID :one
SELECT *
FROM monthly_transactions
WHERE id = $1;

-- name: ListMonthlyTransactionsByUserIDPaginated :many
SELECT 
    id,
    user_id,
    name,
    value,
    day,
    category_id,
    credit_card_id,
    created_at,
    updated_at,
    COUNT(*) OVER() as total_count
FROM monthly_transactions
WHERE user_id = $1
ORDER BY day ASC
LIMIT $2 OFFSET $3;

-- name: UpdateMonthlyTransaction :one
UPDATE monthly_transactions
SET
    name = $2,
    value = $3,
    day = $4,
    category_id = $5,
    credit_card_id = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteMonthlyTransaction :exec
DELETE FROM monthly_transactions
WHERE id = $1;