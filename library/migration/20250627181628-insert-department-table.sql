
-- +migrate Up
INSERT INTO public.departments (id, name, slug, created_at, updated_at)
VALUES 
    (1, 'General Medicine', 'general-medicine', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Cardiology', 'cardiology', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 'Neurology', 'neurology', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 'Pediatrics', 'pediatrics', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 'Orthopedics', 'orthopedics', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +migrate Down
DELETE FROM departments WHERE id IN (1,2,3,4,5);
