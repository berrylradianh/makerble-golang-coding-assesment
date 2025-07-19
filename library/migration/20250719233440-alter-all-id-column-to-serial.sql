
-- +migrate Up

-- Alter table users
CREATE SEQUENCE public.users_id_seq;
ALTER TABLE public.users
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq');
ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
SELECT setval('public.users_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM public.users), false);

-- Alter table roles
CREATE SEQUENCE public.roles_id_seq;
ALTER TABLE public.roles
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq');
ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;
SELECT setval('public.roles_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM public.roles), false);

-- Alter table emergency_contacts
CREATE SEQUENCE public.emergency_contacts_id_seq;
ALTER TABLE public.emergency_contacts
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id SET DEFAULT nextval('public.emergency_contacts_id_seq');
ALTER SEQUENCE public.emergency_contacts_id_seq OWNED BY public.emergency_contacts.id;
SELECT setval('public.emergency_contacts_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM public.emergency_contacts), false);

-- Alter table user_doctors
CREATE SEQUENCE public.user_doctors_id_seq;
ALTER TABLE public.user_doctors
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id SET DEFAULT nextval('public.user_doctors_id_seq');
ALTER SEQUENCE public.user_doctors_id_seq OWNED BY public.user_doctors.id;
SELECT setval('public.user_doctors_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM public.user_doctors), false);

-- Alter table employment_statuses
CREATE SEQUENCE public.employment_statuses_id_seq;
ALTER TABLE public.employment_statuses
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id SET DEFAULT nextval('public.employment_statuses_id_seq');
ALTER SEQUENCE public.employment_statuses_id_seq OWNED BY public.employment_statuses.id;
SELECT setval('public.employment_statuses_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM public.employment_statuses), false);

-- Alter table departments
CREATE SEQUENCE public.departments_id_seq;
ALTER TABLE public.departments
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id SET DEFAULT nextval('public.departments_id_seq');
ALTER SEQUENCE public.departments_id_seq OWNED BY public.departments.id;
SELECT setval('public.departments_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM public.departments), false);

-- Alter table medical_records
CREATE SEQUENCE public.medical_records_id_seq;
ALTER TABLE public.medical_records
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id SET DEFAULT nextval('public.medical_records_id_seq');
ALTER SEQUENCE public.medical_records_id_seq OWNED BY public.medical_records.id;
SELECT setval('public.medical_records_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM public.medical_records), false);

-- +migrate Down

-- Alter table users
ALTER TABLE public.users
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.users_id_seq;

-- Alter table roles
ALTER TABLE public.roles
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.roles_id_seq;

-- Alter table emergency_contacts
ALTER TABLE public.emergency_contacts
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.emergency_contacts_id_seq;

-- Alter table user_doctors
ALTER TABLE public.user_doctors
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.user_doctors_id_seq;

-- Alter table employment_statuses
ALTER TABLE public.employment_statuses
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.employment_statuses_id_seq;

-- Alter table departments
ALTER TABLE public.departments
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.departments_id_seq;

-- Alter table medical_records
ALTER TABLE public.medical_records
    ALTER COLUMN id SET DATA TYPE integer,
    ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.medical_records_id_seq;