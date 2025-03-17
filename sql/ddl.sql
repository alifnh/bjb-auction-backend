DROP TABLE IF EXISTS users_assets;
DROP TABLE IF EXISTS assets;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS category_enum;
DROP TYPE IF EXISTS gender_enum;

-- enum kategori
CREATE TYPE category_enum AS ENUM ('kendaraan', 'properti');

-- enum gender
CREATE TYPE gender_enum AS ENUM ('male', 'female');


CREATE TABLE users (
    id bigserial PRIMARY KEY,
    nik varchar(50) NOT NULL,
    name varchar(50) NOT NULL,
    gender gender_enum NOT NULL,
    date_of_birth Date NOT NULL,
    city varchar(50) NOT NULL,
    email varchar(50) NOT NULL UNIQUE,
    password varchar(255),
    phone_number varchar(13),
    role varchar(50) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);


CREATE TABLE assets (
    id bigserial PRIMARY KEY,
    category category_enum,
    img_url varchar(100) NOT NULL,
    name varchar(50) NOT NULL,
    price double precision,
    description varchar(200),
    city varchar(100) NOT NULL,
    address varchar(100) NOT NULL,
    maps_url varchar(100),
    start_date Date NOT NULL,
    end_date Date NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp,
    updated_at timestamp NOT NULL
);

CREATE TABLE users_assets (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    asset_id bigint NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp,
    updated_at timestamp NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (asset_id) REFERENCES assets(id)
);

CREATE INDEX idx_users_assets_user_id ON users_assets(user_id);
CREATE INDEX idx_users_assets_asset_id ON users_assets(asset_id);
