DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users
(
    user_id      UUID PRIMARY KEY,
    first_name   VARCHAR(32) NOT NULL CHECK (first_name <> ''),
    last_name    VARCHAR(32) NOT NULL CHECK (last_name <> ''),
    email        VARCHAR(64) UNIQUE NOT NULL CHECK (email <> ''),
    password     VARCHAR(250) NOT NULL CHECK (length(password) <> 0),
    created_at   TIMESTAMP NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
    updated_at   TIMESTAMP DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
    login_date   TIMESTAMP NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
);

