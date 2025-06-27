
-- +migrate Up
INSERT INTO public.employment_statuses (id, name, slug, created_at, updated_at)
VALUES 
    (1, 'Permanent', 'permanent', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Contract', 'contract', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +migrate Down
DELETE FROM employment_statuses WHERE id IN (1,2);