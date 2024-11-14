CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE,
    username TEXT UNIQUE,
    password TEXT,
    birthday TEXT,
    gender BOOLEAN,
    preferences TEXT,
    pics TEXT,
    token TEXT,
    location TEXT
);
