CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL PRIMARY KEY,
    full_name VARCHAR(256),
    photo TEXT,
    phone_number VARCHAR(13),
    description TEXT,
);