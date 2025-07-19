
-- +migrate Up
INSERT INTO roles (id, name, slug)
VALUES
        (5, 'Super Admin', 'superadmin');
-- +migrate Down
DELETE FROM roles WHERE id IN (5);