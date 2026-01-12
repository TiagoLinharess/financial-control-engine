-- name: CreateAnnualTransaction :one
INSERT INTO annual_transactions (
    user_id,
    name,
    value,
    day,
    month,
    category_id,
    credit_card_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetAnnualTransactionByID :one
SELECT *
FROM annual_transactions
WHERE id = $1;

-- name: ListAnnualTransactionsByUserIDPaginated :many
SELECT 
    id,
    user_id,
    name,
    value,
    day,
    month,
    category_id,
    credit_card_id,
    created_at,
    updated_at,
    COUNT(*) OVER() as total_count
FROM annual_transactions
WHERE user_id = $1
ORDER BY month ASC, day ASC
LIMIT $2 OFFSET $3;

-- name: UpdateAnnualTransaction :one
UPDATE annual_transactions
SET
    name = $2,
    value = $3,
    day = $4,
    month = $5,
    category_id = $6,
    credit_card_id = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAnnualTransaction :exec
DELETE FROM annual_transactions
WHERE id = $1;