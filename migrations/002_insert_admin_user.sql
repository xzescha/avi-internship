-- migrations/002_insert_admin_user.sql
INSERT INTO users (username, password) VALUES ('admin', 'password')
    ON CONFLICT (username) DO NOTHING;
