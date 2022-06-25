CREATE TABLE posts(
    id INTEGER SERIAL PRIMARY KEY,
    user_id INTEGER,
    title TEXT,
    body TEXT
);