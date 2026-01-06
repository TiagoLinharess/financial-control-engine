-- name: CreateCreditCard :one
INSERT INTO credit_cards (
    user_id, 
    name, 
    first_four_numbers, 
    credit_limit, 
    close_day, 
    expire_day
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListCreditCards :many
SELECT * FROM credit_cards
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetCreditCardByID :one
SELECT * FROM credit_cards
WHERE id = $1;

-- name: UpdateCreditCard :one
UPDATE credit_cards
SET
    name = $2,
    first_four_numbers = $3,
    credit_limit = $4,
    close_day = $5,
    expire_day = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCreditCard :exec
DELETE FROM credit_cards
WHERE id = $1;