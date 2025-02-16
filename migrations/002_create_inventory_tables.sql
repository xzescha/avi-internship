-- migrations/002_create_inventory_tables.sql

-- Создание таблицы для начального баланса пользователей
CREATE TABLE IF NOT EXISTS user_inventory (
                                              user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                                              coins INTEGER NOT NULL DEFAULT 1000
);

-- Создание таблицы для хранения данных о товарах магазина (мерче)
CREATE TABLE IF NOT EXISTS store_items (
                                           id SERIAL PRIMARY KEY,
                                           name TEXT UNIQUE NOT NULL,
                                           price INTEGER NOT NULL
);

-- Вставка товаров магазина (при повторном запуске не дублируется)
INSERT INTO store_items (name, price) VALUES
                                          ('t-shirt', 80),
                                          ('cup', 20),
                                          ('book', 50),
                                          ('pen', 10),
                                          ('powerbank', 200),
                                          ('hoody', 300),
                                          ('umbrella', 200),
                                          ('socks', 10),
                                          ('wallet', 50),
                                          ('pink-hoody', 500)
ON CONFLICT (name) DO NOTHING;
