
-- +migrate Up
INSERT INTO public.users (id, role_id, identity_number, email, password, name, phone, date_of_birth, address, gender, created_at, updated_at)
VALUES 
    (2, 3, 'DOC001', 'doctor1@gmail.com', crypt('password', gen_salt('bf')), 'Dr. John Smith', '+6281234567891', '1975-03-15', '456 Doctor Road, Jakarta', 'Male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 3, 'DOC002', 'doctor2@gmail.com', crypt('password', gen_salt('bf')), 'Dr. Jane Doe', '+6281234567892', '1982-07-22', '789 Medical Lane, Bandung', 'Female', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 3, 'DOC003', 'doctor3@gmail.com', crypt('password', gen_salt('bf')), 'Dr. Michael Brown', '+6281234567893', '1978-11-30', '101 Hospital Ave, Surabaya', 'Male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 3, 'DOC004', 'doctor4@gmail.com', crypt('password', gen_salt('bf')), 'Dr. Sarah Wilson', '+6281234567894', '1985-05-10', '202 Clinic Street, Yogyakarta', 'Female', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (6, 3, 'DOC005', 'doctor5@gmail.com', crypt('password', gen_salt('bf')), 'Dr. David Lee', '+6281234567895', '1970-09-25', '303 Health Road, Bali', 'Male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.user_doctors (id, user_id, employment_status_id, department_id, professional_degree, citizenship, language, gmc_number, specialization, sub_specialization, last_education, year_graduated, professional_certification, employee_number, is_active, created_at, updated_at)
VALUES 
    (1, 2, 1, 1, ARRAY['MD'], 'Indonesian', ARRAY['Indonesian', 'English'], 'GMC001', 'General Medicine', NULL, 'Medical Doctor', '2000', ARRAY['Board Certified'], 'EMP001', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 3, 2, 2, ARRAY['MD', 'PhD'], 'Indonesian', ARRAY['Indonesian', 'English'], 'GMC002', 'Cardiology', 'Interventional Cardiology', 'Doctor of Philosophy', '2005', ARRAY['FACC'], 'EMP002', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 4, 1, 3, ARRAY['MD'], 'Indonesian', ARRAY['Indonesian', 'English'], 'GMC003', 'Neurology', NULL, 'Medical Doctor', '2003', ARRAY['Board Certified'], 'EMP003', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 5, 2, 4, ARRAY['MD'], 'Indonesian', ARRAY['Indonesian', 'English'], 'GMC004', 'Pediatrics', 'Pediatric Oncology', 'Medical Doctor', '2008', ARRAY['Board Certified'], 'EMP004', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (5, 6, 1, 5, ARRAY['MD', 'MSc'], 'Indonesian', ARRAY['Indonesian', 'English'], 'GMC005', 'Orthopedics', 'Sports Medicine', 'Master of Science', '1995', ARRAY['Board Certified'], 'EMP005', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +migrate Down
DELETE FROM users WHERE id IN (2,3,4,5,6);
DELETE FROM user_doctors WHERE id IN (1,2,3,4,5);