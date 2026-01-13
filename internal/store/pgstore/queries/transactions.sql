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
RETURNING id, name, date, value, created_at, updated_at;

-- name: GetTransactionByID :one
SELECT 
    t.id,
    t.user_id, 
    t.name, 
    t.date, 
    t.value, 
    t.created_at, 
    t.updated_at,

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

    mt.id as monthly_transactions_id, 
    mt.day as monthly_transactions_day,

    at.id as annual_transactions_id, 
    at.month as annual_transactions_month, 
    at.day as annual_transactions_day,

    it.id as installment_transactions_id, 
    it.initial_date as installment_transactions_initial_date,  
    it.final_date as installment_transactions_final_date
FROM transactions t
LEFT JOIN categories c ON t.category_id = c.id
LEFT JOIN credit_cards cc ON t.credit_card_id = cc.id
LEFT JOIN monthly_transactions mt ON t.monthly_transactions_id = mt.id
LEFT JOIN annual_transactions at ON t.annual_transactions_id = at.id
LEFT JOIN installment_transactions it ON t.installment_transactions_id = it.id
WHERE t.id = $1;

-- name: ListTransactionsByUserIDPaginated :many
SELECT 
    t.id,
    t.user_id, 
    t.name, 
    t.date, 
    t.value, 
    t.created_at, 
    t.updated_at,

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

    mt.id as monthly_transactions_id, 
    mt.day as monthly_transactions_day,

    at.id as annual_transactions_id, 
    at.month as annual_transactions_month, 
    at.day as annual_transactions_day,

    it.id as installment_transactions_id, 
    it.initial_date as installment_transactions_initial_date,  
    it.final_date as installment_transactions_final_date,

    COUNT(*) OVER() as total_count
FROM transactions t
LEFT JOIN categories c ON t.category_id = c.id
LEFT JOIN credit_cards cc ON t.credit_card_id = cc.id
LEFT JOIN monthly_transactions mt ON t.monthly_transactions_id = mt.id
LEFT JOIN annual_transactions at ON t.annual_transactions_id = at.id
LEFT JOIN installment_transactions it ON t.installment_transactions_id = it.id
WHERE t.user_id = $1
ORDER BY t.date DESC
LIMIT $2 OFFSET $3;

-- name: ListTransactionsByUserAndDate :many
SELECT 
    t.id,
    t.user_id, 
    t.name, 
    t.date, 
    t.value, 
    t.created_at, 
    t.updated_at,

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

    mt.id as monthly_transactions_id, 
    mt.day as monthly_transactions_day,

    at.id as annual_transactions_id, 
    at.month as annual_transactions_month, 
    at.day as annual_transactions_day,

    it.id as installment_transactions_id, 
    it.initial_date as installment_transactions_initial_date,  
    it.final_date as installment_transactions_final_date,

    COUNT(*) OVER() as total_count
FROM transactions t
LEFT JOIN categories c ON t.category_id = c.id
LEFT JOIN credit_cards cc ON t.credit_card_id = cc.id
LEFT JOIN monthly_transactions mt ON t.monthly_transactions_id = mt.id
LEFT JOIN annual_transactions at ON t.annual_transactions_id = at.id
LEFT JOIN installment_transactions it ON t.installment_transactions_id = it.id
WHERE t.user_id = $1
  AND t.date >= $2
  AND t.date <= $3
ORDER BY t.date DESC
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