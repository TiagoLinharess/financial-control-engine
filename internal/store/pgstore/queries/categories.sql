-- name: CreateCategory :one
INSERT INTO categories (
    user_id,
    transaction_type,
    name,
    icon
) VALUES (
    $1, $2, $3, $4
) RETURNING id, user_id, transaction_type, name, icon, created_at, updated_at;