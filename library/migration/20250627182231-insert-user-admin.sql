-- +migrate Up
CREATE EXTENSION IF NOT EXISTS pgcrypto SCHEMA public;

INSERT INTO
    public.users (id, role_id, identity_number, email, password, name, phone, date_of_birth, address, gender, created_at, updated_at)
VALUES
    (1,1,'ADM001','admin@gmail.com', crypt('password', gen_salt('bf')),'Admin User','+6281234567890','1980-01-01','123 Admin Street, Jakarta','Male',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);
-- +migrate Down
DELETE FROM users WHERE id = 1;