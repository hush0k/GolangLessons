CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    email TEXT NOT NULL UNIQUE ,
    name TEXT NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email)

