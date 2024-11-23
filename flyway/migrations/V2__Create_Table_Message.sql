CREATE TABLE message (
    id SERIAL PRIMARY KEY,
    sender_id TEXT,
    receiver_id TEXT,
    timestamp TEXT
);
