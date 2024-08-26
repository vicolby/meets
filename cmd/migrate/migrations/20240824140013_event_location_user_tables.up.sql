CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    address VARCHAR(255),
    latitude REAL,
    longitude REAL,
    city VARCHAR(255),
    country VARCHAR(255)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    first_name VARCHAR(255),
    second_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(255),
    password VARCHAR(255),
    reg_date TIMESTAMPTZ
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    name VARCHAR(255),
    description TEXT,
    date TIMESTAMPTZ,
    location_id INT REFERENCES locations(id),
    organizer_id INT REFERENCES users(id)
);

