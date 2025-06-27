
-- +migrate Up
INSERT INTO roles (id, name, slug)
VALUES
        (1, 'Admin', 'admin'),
        (2, 'Receptionist', 'receptionist'),
        (3, 'Doctor', 'doctor'),
        (4, 'Patient', 'patient');
-- +migrate Down
DELETE FROM roles WHERE id IN (1,2,3,4);