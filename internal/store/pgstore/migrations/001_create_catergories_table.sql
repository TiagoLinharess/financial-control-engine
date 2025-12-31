-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    transaction_type INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Índices para melhorar performance
CREATE INDEX idx_categories_user_id ON categories(user_id);
CREATE INDEX idx_categories_transaction_type ON categories(transaction_type);
CREATE INDEX idx_categories_user_transaction ON categories(user_id, transaction_type);

---- create above / drop below ----

-- Write your migrate down statements here
DROP INDEX IF EXISTS idx_categories_user_transaction;
DROP INDEX IF EXISTS idx_categories_transaction_type;
DROP INDEX IF EXISTS idx_categories_user_id;
DROP TABLE IF EXISTS categories;