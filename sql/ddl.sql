DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name varchar(50) NOT NULL,
    email varchar(50) NOT NULL UNIQUE UNIQUE,
    password varchar(255),
    phone_number varchar(13),
    role varchar(50) NOT NULL,
    city varchar(50) NOT NULL,
    gender varchar(20) NOT NULL,
    date_of_birth date NOT NULL,
    nik varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP NOT NULL default current_timestamp,
    deleted_at TIMESTAMP
);