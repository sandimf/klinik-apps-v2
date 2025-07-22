-- Tabel users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(32) NOT NULL
);

-- Tabel patients
CREATE TABLE patients (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    nik VARCHAR(32),
    full_name VARCHAR(255) NOT NULL,
    birth_place VARCHAR(128),
    birth_date VARCHAR(32),
    gender VARCHAR(16),
    address TEXT,
    rt VARCHAR(8),
    rw VARCHAR(8),
    village VARCHAR(128),
    district VARCHAR(128),
    religion VARCHAR(32),
    marital VARCHAR(32),
    job VARCHAR(64),
    nationality VARCHAR(64),
    valid_until VARCHAR(32),
    blood_type VARCHAR(8),
    height INT,
    weight INT,
    age INT,
    email VARCHAR(255),
    phone VARCHAR(32),
    ktp_images TEXT[],
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Tabel doctors
CREATE TABLE doctors (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    full_name VARCHAR(255) NOT NULL,
    nik VARCHAR(32),
    phone_number VARCHAR(32),
    address TEXT,
    specialty VARCHAR(128),
    license_number VARCHAR(64)
);

-- Tabel paramedics
CREATE TABLE paramedics (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    full_name VARCHAR(255) NOT NULL,
    nik VARCHAR(32),
    phone_number VARCHAR(32),
    address TEXT
);

-- Tabel admins
CREATE TABLE admins (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    full_name VARCHAR(255) NOT NULL,
    nik VARCHAR(32),
    phone_number VARCHAR(32),
    address TEXT
);

-- Tabel cashiers
CREATE TABLE cashiers (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    full_name VARCHAR(255) NOT NULL,
    nik VARCHAR(32),
    phone_number VARCHAR(32),
    address TEXT
);

-- Tabel medicines
CREATE TABLE medicines (
    id UUID PRIMARY KEY,
    barcode VARCHAR(64),
    medicine_name VARCHAR(255) NOT NULL,
    brand_name VARCHAR(128),
    category VARCHAR(64),
    dosage INT,
    content TEXT,
    quantity INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Tabel medical_records
CREATE TABLE medical_records (
    id UUID PRIMARY KEY,
    patient_id UUID REFERENCES patients(id),
    mr_number VARCHAR(32) NOT NULL,
    created_at TIMESTAMP
);

-- Tabel physical_examinations
CREATE TABLE physical_examinations (
    id UUID PRIMARY KEY,
    patient_id UUID REFERENCES patients(id),
    paramedis_id UUID REFERENCES paramedics(id),
    doctor_id UUID REFERENCES doctors(id),
    blood_pressure VARCHAR(32),
    heart_rate INT,
    oxygen_saturation INT,
    respiratory_rate INT,
    body_temperature FLOAT,
    physical_assessment TEXT,
    reason TEXT,
    medical_advice TEXT,
    health_status VARCHAR(64),
    pendampingan TEXT[],
    konsultasi_dokter BOOLEAN,
    konsultasi_dokter_status VARCHAR(32),
    doctor_advice TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Tabel screening_questions
CREATE TABLE screening_questions (
    id UUID PRIMARY KEY,
    label TEXT NOT NULL,
    type VARCHAR(32) NOT NULL,
    options TEXT[]
);

-- Tabel screening_answers
CREATE TABLE screening_answers (
    id UUID PRIMARY KEY,
    patient_info JSONB NOT NULL,
    answers JSONB NOT NULL,
    created_at TIMESTAMP
);

-- Tabel screening_queues
CREATE TABLE screening_queues (
    id UUID PRIMARY KEY,
    patient_info JSONB NOT NULL,
    screening_answer_id UUID REFERENCES screening_answers(id),
    status VARCHAR(32) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
); 