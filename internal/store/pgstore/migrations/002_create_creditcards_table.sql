-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS credit_cards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    first_four_numbers VARCHAR(4) NOT NULL,
    credit_limit DOUBLE PRECISION NOT NULL,
    close_day INTEGER NOT NULL,
    expire_day INTEGER NOT NULL,
    background_color VARCHAR(10) NOT NULL,
    text_color VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Índices para melhorar performance

CREATE INDEX idx_credit_cards_user_id ON credit_cards(user_id);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible

DROP INDEX IF EXISTS idx_credit_cards_user_id;
DROP TABLE IF EXISTS credit_cards;

-- Then delete the separator line above.
