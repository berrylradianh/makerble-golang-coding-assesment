-- +migrate Up
CREATE EXTENSION IF NOT EXISTS pgcrypto SCHEMA public;

INSERT INTO
    public.users (id, role_id, identity_number, email, password, name, phone, date_of_birth, address, gender, created_at, updated_at)
VALUES
    (7,2,'RCP001','receptionist1@gmail.com', crypt('password', gen_salt('bf')),'Receptionist','+628453215468875','1980-01-01','123 Admin Street, Jakarta','Male',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);  
-- +migrate Down
DELETE FROM users WHERE id = 7;