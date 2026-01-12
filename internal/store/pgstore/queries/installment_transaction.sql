-- name: CreateInstallmentTransaction :one
INSERT INTO installment_transactions (
    user_id,
    name,
    value,
    initial_date,
    final_date,
    category_id,
    credit_card_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetInstallmentTransactionByID :one
SELECT *
FROM installment_transactions
WHERE id = $1;

-- name: ListInstallmentTransactionsByUserIDPaginated :many
SELECT 
    id,
    user_id,
    name,
    value,
    initial_date,
    final_date,
    category_id,
    credit_card_id,
    created_at,
    updated_at,
    COUNT(*) OVER() as total_count
FROM installment_transactions
WHERE user_id = $1
ORDER BY initial_date DESC
LIMIT $2 OFFSET $3;

-- name: UpdateInstallmentTransaction :one
UPDATE installment_transactions
SET
    name = $2,
    value = $3,
    initial_date = $4,
    final_date = $5,
    category_id = $6,
    credit_card_id = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteInstallmentTransaction :exec
DELETE FROM installment_transactions
WHERE id = $1;