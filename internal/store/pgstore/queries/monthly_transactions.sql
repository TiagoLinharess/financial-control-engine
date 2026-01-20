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
SELECT 
    mt.id,
    mt.user_id, 
    mt.name, 
    mt.value,
    mt.day, 
    mt.created_at, 
    mt.updated_at,

    c.id as category_id, 
    c.transaction_type as category_transaction_type, 
    c.name as category_name, 
    c.icon as category_icon,

    cc.id as creditcard_id, 
    cc.name as creditcard_name, 
    cc.first_four_numbers as creditcard_first_four_numbers, 
    cc.credit_limit as creditcard_credit_limit, 
    cc.close_day as creditcard_close_day, 
    cc.expire_day as creditcard_expire_day, 
    cc.background_color as creditcard_background_color, 
    cc.text_color as creditcard_text_color,

    COUNT(*) OVER() as total_count
FROM monthly_transactions mt
LEFT JOIN categories c ON mt.category_id = c.id
LEFT JOIN credit_cards cc ON mt.credit_card_id = cc.id
WHERE mt.id = $1;

-- name: ListMonthlyTransactionsByUserIDPaginated :many
SELECT 
    mt.id,
    mt.user_id, 
    mt.name, 
    mt.value,
    mt.day, 
    mt.created_at, 
    mt.updated_at,

    c.id as category_id, 
    c.transaction_type as category_transaction_type, 
    c.name as category_name, 
    c.icon as category_icon,

    cc.id as creditcard_id, 
    cc.name as creditcard_name, 
    cc.first_four_numbers as creditcard_first_four_numbers, 
    cc.credit_limit as creditcard_credit_limit, 
    cc.close_day as creditcard_close_day, 
    cc.expire_day as creditcard_expire_day, 
    cc.background_color as creditcard_background_color, 
    cc.text_color as creditcard_text_color,

    COUNT(*) OVER() as total_count
FROM monthly_transactions mt
LEFT JOIN categories c ON mt.category_id = c.id
LEFT JOIN credit_cards cc ON mt.credit_card_id = cc.id
WHERE mt.user_id = $1
ORDER BY mt.day ASC
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