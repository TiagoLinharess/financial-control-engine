-- Write your migrate up statements here

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    value NUMERIC(15, 2) NOT NULL,
    category_id UUID NOT NULL,
    credit_card_id UUID,
    monthly_transactions_id UUID,
    annual_transactions_id UUID,
    installment_transactions_id UUID,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT,
    CONSTRAINT fk_credit_card FOREIGN KEY (credit_card_id) REFERENCES credit_cards(id) ON DELETE SET NULL,
    CONSTRAINT fk_monthly_transactions FOREIGN KEY (monthly_transactions_id) REFERENCES monthly_transactions(id) ON DELETE SET NULL,
    CONSTRAINT fk_annual_transactions FOREIGN KEY (annual_transactions_id) REFERENCES annual_transactions(id) ON DELETE SET NULL,
    CONSTRAINT fk_installment_transactions FOREIGN KEY (installment_transactions_id) REFERENCES installment_transactions(id) ON DELETE SET NULL
);

-- Índices para melhor performance
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_date ON transactions(date);
CREATE INDEX idx_transactions_category ON transactions(category_id);
CREATE INDEX idx_transactions_credit_card ON transactions(credit_card_id);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

DROP INDEX IF EXISTS idx_transactions_credit_card;
DROP INDEX IF EXISTS idx_transactions_category;
DROP INDEX IF EXISTS idx_transactions_date;
DROP INDEX IF EXISTS idx_transactions_user_id;
DROP TABLE IF EXISTS transactions;