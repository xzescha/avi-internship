-- migrations/001_create_users_table.sql
-- migrations/001_create_users_table.sql

-- Включаем расширение для генерации UUID, если ещё не включено
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создание таблицы пользователей с UUID в качестве первичного ключа
CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     username TEXT UNIQUE NOT NULL,
                                     password TEXT NOT NULL,
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
