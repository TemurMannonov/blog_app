CREATE TABLE IF NOT EXISTS blogs(
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT NOT NULL,
    author VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);