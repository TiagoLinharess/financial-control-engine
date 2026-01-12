-- Write your migrate up statements here

CREATE TABLE monthly_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    value NUMERIC(15, 2) NOT NULL,
    day INTEGER NOT NULL,
    category_id UUID NOT NULL,
    credit_card_id UUID,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT,
    CONSTRAINT fk_credit_card FOREIGN KEY (credit_card_id) REFERENCES credit_cards(id) ON DELETE SET NULL
);

-- Índices para melhor performance
CREATE INDEX idx_monthly_transactions_user_id ON monthly_transactions(user_id);
CREATE INDEX idx_monthly_transactions_category ON monthly_transactions(category_id);
CREATE INDEX idx_monthly_transactions_credit_card ON monthly_transactions(credit_card_id);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

DROP INDEX IF EXISTS idx_monthly_transactions_credit_card;
DROP INDEX IF EXISTS idx_monthly_transactions_category;
DROP INDEX IF EXISTS idx_monthly_transactions_date;
DROP INDEX IF EXISTS idx_monthly_transactions_user_id;
DROP TABLE IF EXISTS monthly_transactions;