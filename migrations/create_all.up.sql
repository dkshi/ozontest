CREATE TABLE urls (
    url_id SERIAL PRIMARY KEY,
    original_url VARCHAR(300),
    short_url VARCHAR(50),
    ttl TIMESTAMP
);