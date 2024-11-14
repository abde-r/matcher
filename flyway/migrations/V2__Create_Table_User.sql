CREATE TABLE users (
    id SERIAL PRIMARY KEY,                -- Auto-incremented ID
    first_name VARCHAR(100) NOT NULL,     -- First name
    last_name VARCHAR(100) NOT NULL,      -- Last name
    email VARCHAR(255) UNIQUE NOT NULL,   -- Email, must be unique
    username VARCHAR(50) UNIQUE NOT NULL, -- Username, must be unique
    password VARCHAR(255) NOT NULL,       -- Password
    birthday DATE NOT NULL,               -- Birthday (as date)
    gender BOOLEAN NOT NULL,              -- Gender (assuming TRUE for male, FALSE for female)
    preferences TEXT,                     -- Preferences as a string (could be JSON if more complex)
    pics TEXT,                            -- Pics (URL or JSON of image paths)
    token VARCHAR(255),                   -- Token (assuming it's for authentication or session)
    location VARCHAR(255)                 -- Location of the user
);
