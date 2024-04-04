CREATE TABLE urls {
    url_id SERIAL PRIMARY KEY,
    short_url VARCHAR(50),
    original_url VARCHAR(300)
};