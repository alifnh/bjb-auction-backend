DROP TABLE IF EXISTS users_assets;
DROP TABLE IF EXISTS assets;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS category_enum;

-- enum kategori
CREATE TYPE category_enum AS ENUM ('kendaraan', 'properti');



CREATE TABLE users (
    id bigserial PRIMARY KEY,
    nik varchar(50) NOT NULL,
    name varchar(50) NOT NULL,
    gender varchar(20) NOT NULL,
    date_of_birth Date NOT NULL,
    city varchar(50) NOT NULL,
    email varchar(50) NOT NULL UNIQUE,
    password varchar(255),
    phone_number varchar(13),
    role varchar(50) NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp,
    deleted_at timestamp
);


CREATE TABLE assets (
    id bigserial PRIMARY KEY,
    category category_enum,
    img_url varchar(100) NOT NULL,
    name varchar(50) NOT NULL,
    price double precision,
    description varchar(500),
    city varchar(100) NOT NULL,
    address varchar(500) NOT NULL,
    maps_url varchar(500),
    start_date Date NOT NULL,
    end_date Date NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    deleted_at timestamp,
    updated_at timestamp NOT NULL default current_timestamp

);

CREATE TABLE users_assets (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    asset_id bigint NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    deleted_at timestamp,
    updated_at timestamp NOT NULL default current_timestamp,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (asset_id) REFERENCES assets(id)
);