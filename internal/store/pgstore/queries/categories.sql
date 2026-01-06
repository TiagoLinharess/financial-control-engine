-- name: CreateCategory :one
INSERT INTO categories (
    user_id,
    transaction_type,
    name,
    icon
) VALUES (
    $1, $2, $3, $4
) RETURNING id, user_id, transaction_type, name, icon, created_at, updated_at;

-- name: GetCategoriesByUserID :many
SELECT id, user_id, transaction_type, name, icon, created_at, updated_at
FROM categories
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CountCategoriesByUserID :one
SELECT COUNT(*) as total
FROM categories
WHERE user_id = $1;

-- name: GetCategoryByID :one
SELECT id, user_id, transaction_type, name, icon, created_at, updated_at
FROM categories
WHERE id = $1;

-- name: UpdateCategory :one
UPDATE categories
SET 
    name = $2,
    icon = $3,
    transaction_type = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, transaction_type, name, icon, created_at, updated_at;

-- name: DeleteCategoryByID :exec
DELETE FROM categories
WHERE id = $1;