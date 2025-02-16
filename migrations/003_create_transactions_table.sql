-- migrations/003_create_transactions_table.sql

-- Создание таблицы для хранения истории транзакций по всем пользователям
CREATE TABLE IF NOT EXISTS transactions (
                                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    transaction_type TEXT NOT NULL, -- например, 'credit', 'debit', 'purchase'
    amount INTEGER NOT NULL,        -- сумма изменения (положительное или отрицательное значение)
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );
