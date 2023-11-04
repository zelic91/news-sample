-- migrate:up
CREATE TABLE admins (
    id BIGSERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    username TEXT NOT NULL,
    email TEXT,
    password_hashed TEXT,
    password_salt TEXT,
    status TEXT,
    created_by BIGSERIAL,
    updated_by BIGSERIAL,
    deleted_by BIGSERIAL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_admins_unique_username ON admins (username);
CREATE UNIQUE INDEX idx_admins_unique_email ON admins (email);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    username TEXT NOT NULL,
    email TEXT,
    password_hashed TEXT,
    password_salt TEXT,
    status TEXT,
    created_by BIGSERIAL,
    updated_by BIGSERIAL,
    deleted_by BIGSERIAL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_users_unique_username ON users (username);
CREATE UNIQUE INDEX idx_users_unique_email ON users (email);
CREATE TABLE devices (
    id BIGSERIAL PRIMARY KEY,
    device_token VARCHAR,
    platform VARCHAR,
    status VARCHAR,

    created_by BIGSERIAL,
    updated_by BIGSERIAL,
    deleted_by BIGSERIAL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_created_by_user
        FOREIGN KEY(created_by)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_updated_by_user
        FOREIGN KEY(created_by)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_deleted_by_user
        FOREIGN KEY(deleted_by)
            REFERENCES users(id)
            ON DELETE CASCADE
);
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    description VARCHAR,
    title VARCHAR,

    created_by BIGSERIAL,
    updated_by BIGSERIAL,
    deleted_by BIGSERIAL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE objects (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGSERIAL NOT NULL ,
    content VARCHAR,
    title VARCHAR,

    created_by BIGSERIAL,
    updated_by BIGSERIAL,
    deleted_by BIGSERIAL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
        
    CONSTRAINT fk_category_id
        FOREIGN KEY(category_id)
            REFERENCES categories(id)
            ON DELETE CASCADE
);


-- migrate:down
DROP TABLE devices;
DROP TABLE categories;
DROP TABLE objects;

DROP TABLE users;
DROP TABLE admins;
