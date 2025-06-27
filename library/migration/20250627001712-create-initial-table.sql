
-- +migrate Up
CREATE TABLE public.users
(
    id integer NOT NULL,
    code uuid NOT NULL DEFAULT gen_random_uuid(),
    role_id integer NOT NULL,
    identity_number character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    phone character varying(20) NOT NULL,
    date_of_birth date NOT NULL,
    address text NOT NULL,
    gender character varying(20) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_user PRIMARY KEY (id)
        INCLUDE(id),
    CONSTRAINT uq_email UNIQUE (email)
        INCLUDE(email),
    CONSTRAINT uq_phone UNIQUE (phone)
        INCLUDE(phone),
    CONSTRAINT uq_identity_number UNIQUE (identity_number)
        INCLUDE(identity_number)
);

CREATE TABLE public.roles
(
    id integer NOT NULL,
    code uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(50) NOT NULL,
    slug character varying(50) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_roles PRIMARY KEY (id)
);

CREATE TABLE public.emergency_contacts
(
    id integer NOT NULL,
    user_id integer NOT NULL,
    code uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(255) NOT NULL,
    phone character varying(20) NOT NULL,
    relationship character varying(50) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT uq_phont UNIQUE (phone)
        INCLUDE(phone)
);

CREATE TABLE public.user_doctors
(
    id integer NOT NULL,
    code uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id integer NOT NULL,
    employment_status_id integer NOT NULL,
    department_id integer NOT NULL,
    professional_degree character varying(10)[] NOT NULL,
    citizenship character varying(255) NOT NULL,
    language character varying(20)[] NOT NULL,
    gmc_number character varying(255) NOT NULL,
    specialization character varying(255),
    sub_specialization character varying(255),
    last_education character varying(255) NOT NULL,
    year_graduated character varying(10) NOT NULL,
    professional_certification character varying(255)[],
    employee_number character varying(255) NOT NULL,
    is_active boolean NOT NULL DEFAULT false,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE public.employment_statuses
(
    id integer NOT NULL,
    code uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE public.departments
(
    id integer NOT NULL,
    code uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE public.medical_records
(
    id integer NOT NULL,
    code uuid NOT NULL,
    user_id integer NOT NULL,
    doctor_id integer NOT NULL,
    current_complaint text NOT NULL,
    disease_history character varying(255)[],
    medicine_allergy character varying(255)[],
    medication_taken character varying(255)[],
    is_ever_surgery boolean NOT NULL DEFAULT false,
    assigned_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE public.users
    ADD CONSTRAINT fk_user_role FOREIGN KEY (role_id)
    REFERENCES public.roles (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE public.emergency_contacts
    ADD CONSTRAINT fk_emergency_contact_user FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE public.user_doctors
    ADD CONSTRAINT fk_user_doctor_user FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE public.user_doctors
    ADD CONSTRAINT fk_user_doctor_employment_status FOREIGN KEY (employment_status_id)
    REFERENCES public.employment_statuses (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE public.user_doctors
    ADD CONSTRAINT fk_user_doctor_department FOREIGN KEY (department_id)
    REFERENCES public.departments (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE public.medical_records
    ADD CONSTRAINT fk_medical_record_user FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE public.medical_records
    ADD CONSTRAINT fk_medical_record_user_doctor FOREIGN KEY (doctor_id)
    REFERENCES public.user_doctors (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
    
-- +migrate Down

ALTER TABLE IF EXISTS public.medical_records
    DROP CONSTRAINT IF EXISTS fk_medical_record_user;

ALTER TABLE IF EXISTS public.medical_records
    DROP CONSTRAINT IF EXISTS fk_medical_record_user_doctor;

ALTER TABLE IF EXISTS public.user_doctors
    DROP CONSTRAINT IF EXISTS fk_user_doctor_user;

ALTER TABLE IF EXISTS public.user_doctors
    DROP CONSTRAINT IF EXISTS fk_user_doctor_employment_status;

ALTER TABLE IF EXISTS public.user_doctors
    DROP CONSTRAINT IF EXISTS fk_user_doctor_department;

ALTER TABLE IF EXISTS public.emergency_contacts
    DROP CONSTRAINT IF EXISTS fk_emergency_contact_user;

ALTER TABLE IF EXISTS public.users
    DROP CONSTRAINT IF EXISTS fk_user_role;

DROP TABLE IF EXISTS public.medical_records;
DROP TABLE IF EXISTS public.user_doctors;
DROP TABLE IF EXISTS public.emergency_contacts;
DROP TABLE IF EXISTS public.departments;
DROP TABLE IF EXISTS public.employment_statuses;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.roles;